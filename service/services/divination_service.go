package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/hobbyqhd/yijing/service/config"
	"github.com/hobbyqhd/yijing/service/models"
	gopenai "github.com/sashabaranov/go-openai"
)

type DivinationService struct{}

type DivinationRequest struct {
	Type     string      `json:"type" binding:"required"`
	Question string      `json:"question" binding:"required"`
	Input    interface{} `json:"input,omitempty"`
}

func NewDivinationService() *DivinationService {
	return &DivinationService{}
}

func (s *DivinationService) CreateDivination(userId uint, req *DivinationRequest) (*models.Divination, error) {
	// 根据占卜类型处理输入数据
	var result interface{}
	var err error

	if !models.IsValidDivinationType(req.Type) {
		return nil, fmt.Errorf("不支持的占卜类型")
	}

	switch models.DivinationType(req.Type) {
	case models.TypeTarot:
		result, err = s.drawTarotCards()
	case models.TypeBazi:
		if birthTime, ok := req.Input.(string); ok {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", birthTime)
			if err != nil {
				return nil, fmt.Errorf("出生时间格式错误")
			}
			result, err = s.generateBaziChart(parsedTime)
		} else {
			return nil, fmt.Errorf("八字占卜需要提供出生时间")
		}
	case models.TypeZodiac:
		// 处理星座占卜
		if zodiacSign, ok := req.Input.(string); ok {
			result = map[string]interface{}{
				"sign": zodiacSign,
				"date": time.Now().Format("2006-01-02"),
			}
		} else {
			return nil, fmt.Errorf("星座占卜需要提供星座信息")
		}
	case models.TypeYijing:
		// 简单实现：随机生成卦象
		result = map[string]interface{}{
			"hexagram":       rand.Intn(64) + 1,
			"changing_lines": []int{rand.Intn(6) + 1},
		}
	default:
		return nil, fmt.Errorf("不支持的占卜类型")
	}

	if err != nil {
		return nil, err
	}

	// 获取AI解析
	analysis, err := s.getAIAnalysis(req)
	if err != nil {
		return nil, err
	}

	// 将result转换为JSON字符串
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("结果序列化失败: %v", err)
	}

	// 创建占卜记录
	divination := &models.Divination{
		UserID:     userId,
		Type:       models.DivinationType(req.Type),
		Question:   req.Question,
		Input:      string(resultJSON),
		Result:     string(resultJSON),
		AIAnalysis: analysis,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// 保存到数据库
	if err := config.DB.Create(divination).Error; err != nil {
		return nil, fmt.Errorf("保存占卜记录失败: %v", err)
	}

	return divination, nil
}

func (s *DivinationService) GetUserDivinations(userId uint) ([]*models.Divination, error) {
	var divinations []*models.Divination
	if err := config.DB.Where("user_id = ?", userId).Find(&divinations).Error; err != nil {
		return nil, fmt.Errorf("获取占卜历史记录失败: %v", err)
	}
	return divinations, nil
}

// drawTarotCards 抽取塔罗牌
func (s *DivinationService) drawTarotCards() (*models.TarotReading, error) {
	// 定义三张牌阵
	spread := models.TarotSpread{
		ID:          1,
		Name:        "三张牌阵",
		Description: "过去、现在、未来",
		Positions:   3,
	}

	// 创建一个包含所有塔罗牌的切片
	allCards := make([]models.TarotCard, 78)
	for i := 0; i < 22; i++ { // 大阿卡纳
		allCards[i] = models.TarotCard{
			ID:     i,
			Type:   "major",
			Number: i,
		}
	}
	for i := 0; i < 56; i++ { // 小阿卡纳
		suits := []string{"wands", "cups", "swords", "pentacles"}
		allCards[i+22] = models.TarotCard{
			ID:     i + 22,
			Type:   "minor",
			Suit:   suits[i/14],
			Number: i%14 + 1,
		}
	}

	// 随机抽取指定数量的牌
	selectedCards := make([]models.TarotCard, spread.Positions)
	positions := make([]bool, spread.Positions)
	for i := 0; i < spread.Positions; i++ {
		// 随机选择一张牌
		cardIndex := rand.Intn(len(allCards))
		selectedCards[i] = allCards[cardIndex]
		// 移除已选择的牌
		allCards = append(allCards[:cardIndex], allCards[cardIndex+1:]...)
		// 随机决定正逆位
		positions[i] = rand.Intn(2) == 1
	}

	return &models.TarotReading{
		Cards:     selectedCards,
		Positions: positions,
		Spread:    spread,
	}, nil
}

// generateBaziChart 生成八字命盘
func (s *DivinationService) generateBaziChart(birthTime time.Time) (*models.BaziReading, error) {
	// 计算年柱
	yearStem := s.calculateYearStem(birthTime.Year())
	yearBranch := s.calculateYearBranch(birthTime.Year())

	// 计算月柱
	monthStem := s.calculateMonthStem(yearStem, birthTime.Month())
	monthBranch := s.calculateMonthBranch(birthTime.Month())

	// 计算日柱
	dayStem := s.calculateDayStem(birthTime)
	dayBranch := s.calculateDayBranch(birthTime)

	// 计算时柱
	hourStem := s.calculateHourStem(dayStem, birthTime.Hour())
	hourBranch := s.calculateHourBranch(birthTime.Hour())

	// 构建八字命盘
	chart := models.BaziChart{
		Year: models.BaziPillar{
			Stem:   yearStem,
			Branch: yearBranch,
		},
		Month: models.BaziPillar{
			Stem:   monthStem,
			Branch: monthBranch,
		},
		Day: models.BaziPillar{
			Stem:   dayStem,
			Branch: dayBranch,
		},
		Hour: models.BaziPillar{
			Stem:   hourStem,
			Branch: hourBranch,
		},
	}

	// 计算五行分布
	elements := s.calculateElements(chart)

	return &models.BaziReading{
		Chart:     chart,
		DayMaster: dayStem,
		Elements:  elements,
	}, nil
}

// calculateYearStem 计算年干
func (s *DivinationService) calculateYearStem(year int) models.Stem {
	stems := []models.Stem{models.Gui, models.Jia, models.Yi, models.Bing, models.Ding,
		models.Wu4, models.Ji, models.Geng, models.Xin, models.Ren}
	return stems[year%10]
}

// calculateYearBranch 计算年支
func (s *DivinationService) calculateYearBranch(year int) models.Branch {
	branches := []models.Branch{models.Hai, models.Zi, models.Chou, models.Yin, models.Mao,
		models.Chen, models.Si, models.Wu, models.Wei, models.Shen, models.You, models.Xu}
	return branches[year%12]
}

// calculateMonthStem 计算月干
func (s *DivinationService) calculateMonthStem(yearStem models.Stem, month time.Month) models.Stem {
	// 根据年干确定月干起始
	baseIndex := map[models.Stem]int{
		models.Jia: 0, models.Yi: 2, models.Bing: 4, models.Ding: 6,
		models.Wu4: 8, models.Ji: 0, models.Geng: 2, models.Xin: 4,
		models.Ren: 6, models.Gui: 8,
	}
	stems := []models.Stem{models.Jia, models.Yi, models.Bing, models.Ding, models.Wu4,
		models.Ji, models.Geng, models.Xin, models.Ren, models.Gui}
	index := (baseIndex[yearStem] + int(month) - 1) % 10
	return stems[index]
}

// calculateMonthBranch 计算月支
func (s *DivinationService) calculateMonthBranch(month time.Month) models.Branch {
	branches := []models.Branch{models.Yin, models.Mao, models.Chen, models.Si,
		models.Wu, models.Wei, models.Shen, models.You, models.Xu, models.Hai,
		models.Zi, models.Chou}
	return branches[month-1]
}

// calculateDayStem 计算日干
func (s *DivinationService) calculateDayStem(birthTime time.Time) models.Stem {
	// 使用简化算法，实际应考虑更复杂的历法计算
	stems := []models.Stem{models.Jia, models.Yi, models.Bing, models.Ding, models.Wu4,
		models.Ji, models.Geng, models.Xin, models.Ren, models.Gui}
	dayNum := birthTime.YearDay()
	return stems[dayNum%10]
}

// calculateDayBranch 计算日支
func (s *DivinationService) calculateDayBranch(birthTime time.Time) models.Branch {
	// 使用简化算法，实际应考虑更复杂的历法计算
	branches := []models.Branch{models.Zi, models.Chou, models.Yin, models.Mao,
		models.Chen, models.Si, models.Wu, models.Wei, models.Shen, models.You,
		models.Xu, models.Hai}
	dayNum := birthTime.YearDay()
	return branches[dayNum%12]
}

// calculateHourStem 计算时干
func (s *DivinationService) calculateHourStem(dayStem models.Stem, hour int) models.Stem {
	// 根据日干确定时干起始
	baseIndex := map[models.Stem]int{
		models.Jia: 0, models.Yi: 2, models.Bing: 4, models.Ding: 6,
		models.Wu4: 8, models.Ji: 0, models.Geng: 2, models.Xin: 4,
		models.Ren: 6, models.Gui: 8,
	}
	stems := []models.Stem{models.Jia, models.Yi, models.Bing, models.Ding, models.Wu4,
		models.Ji, models.Geng, models.Xin, models.Ren, models.Gui}
	index := (baseIndex[dayStem] + hour/2) % 10
	return stems[index]
}

// calculateHourBranch 计算时支
func (s *DivinationService) calculateHourBranch(hour int) models.Branch {
	branches := []models.Branch{models.Zi, models.Chou, models.Yin, models.Mao,
		models.Chen, models.Si, models.Wu, models.Wei, models.Shen, models.You,
		models.Xu, models.Hai}
	return branches[(hour/2)%12]
}

// calculateElements 计算五行分布
func (s *DivinationService) calculateElements(chart models.BaziChart) []models.Element {
	elements := make([]models.Element, 0)
	// 简化实现：仅返回基本五行
	elements = append(elements,
		models.Wood,
		models.Fire,
		models.Earth,
		models.Metal,
		models.Water,
	)
	return elements
}

// getAIAnalysis 获取AI解析
func (s *DivinationService) getAIAnalysis(req *DivinationRequest) (string, error) {
	// 构建提示信息
	prompt := fmt.Sprintf(
		"请根据以下占卜信息进行分析：\n"+
			"占卜类型：%s\n"+
			"问题：%s\n"+
			"请给出详细的解析和建议。",
		req.Type,
		req.Question,
	)

	// 调用OpenAI API
	resp, err := config.OpenAIClient.CreateChatCompletion(
		context.Background(),
		gopenai.ChatCompletionRequest{
			Model: gopenai.GPT3Dot5Turbo,
			Messages: []gopenai.ChatCompletionMessage{
				{
					Role:    gopenai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
