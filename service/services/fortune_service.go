package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/hobbyqhd/yijing/service/config"
	"github.com/hobbyqhd/yijing/service/models"
	gopenai "github.com/sashabaranov/go-openai"
)

type FortuneService struct{}

func NewFortuneService() *FortuneService {
	return &FortuneService{}
}

// CalculateFortune 计算用户运势
func (s *FortuneService) CalculateFortune(userId uint) (*models.Fortune, error) {
	// 检查今日是否已经计算过运势
	today := time.Now().Truncate(24 * time.Hour)
	var existingFortune models.Fortune
	result := config.DB.Where("user_id = ? AND date = ?", userId, today).First(&existingFortune)
	if result.RowsAffected > 0 {
		return &existingFortune, nil
	}

	// 生成运势指数
	fortune := &models.Fortune{
		UserID:       userId,
		Date:         today,
		OverallScore: s.generateScore(),
		LoveScore:    s.generateScore(),
		CareerScore:  s.generateScore(),
		HealthScore:  s.generateScore(),
		WealthScore:  s.generateScore(),
	}

	// 生成AI分析报告
	analysis, suggestions, err := s.generateAIAnalysis(fortune)
	if err != nil {
		return nil, fmt.Errorf("生成AI分析报告失败: %v", err)
	}

	fortune.Analysis = analysis
	fortune.Suggestions = suggestions

	// 保存到数据库
	if err := config.DB.Create(fortune).Error; err != nil {
		return nil, fmt.Errorf("保存运势记录失败: %v", err)
	}

	return fortune, nil
}

// GetUserFortunes 获取用户的运势历史记录
func (s *FortuneService) GetUserFortunes(userId uint, startDate, endDate time.Time) ([]models.Fortune, error) {
	var fortunes []models.Fortune
	result := config.DB.Where("user_id = ? AND date BETWEEN ? AND ?", userId, startDate, endDate).
		Order("date desc").Find(&fortunes)
	return fortunes, result.Error
}

// generateScore 生成0-100之间的运势指数
func (s *FortuneService) generateScore() int {
	return rand.Intn(101)
}

// generateAIAnalysis 生成AI分析报告和建议
func (s *FortuneService) generateAIAnalysis(fortune *models.Fortune) (string, string, error) {
	// 构建提示信息
	prompt := fmt.Sprintf(
		"请根据以下运势指数进行分析和给出建议：\n"+
			"总体运势：%d\n"+
			"感情运势：%d\n"+
			"事业运势：%d\n"+
			"健康运势：%d\n"+
			"财运指数：%d\n"+
			"请分别给出详细的运势分析和具体的建议。",
		fortune.OverallScore,
		fortune.LoveScore,
		fortune.CareerScore,
		fortune.HealthScore,
		fortune.WealthScore,
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
		return "", "", err
	}

	// 解析AI响应，分离分析和建议
	response := resp.Choices[0].Message.Content
	// 这里可以添加更复杂的解析逻辑，目前简单返回相同的内容
	return response, response, nil
}
