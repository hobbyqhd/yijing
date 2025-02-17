package models

// TarotCard 塔罗牌模型
type TarotCard struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`        // 牌名
	NameEn      string `json:"name_en"`     // 英文名
	Type        string `json:"type"`        // 大阿卡纳或小阿卡纳
	Suit        string `json:"suit"`        // 如果是小阿卡纳，表示所属花色
	Number      int    `json:"number"`      // 牌号
	Description string `json:"description"` // 牌面描述
	Upright     string `json:"upright"`     // 正位含义
	Reversed    string `json:"reversed"`    // 逆位含义
	ImageURL    string `json:"image_url"`   // 牌面图片URL
}

// TarotSpread 塔罗牌阵模型
type TarotSpread struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`        // 牌阵名称
	Description string `json:"description"` // 牌阵描述
	Positions   int    `json:"positions"`   // 所需卡牌数量
}

// TarotReading 塔罗牌占卜结果
type TarotReading struct {
	Cards     []TarotCard `json:"cards"`     // 抽取的卡牌
	Positions []bool      `json:"positions"` // 每张牌是否正位
	Spread    TarotSpread `json:"spread"`    // 使用的牌阵
}
