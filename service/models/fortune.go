package models

import (
	"time"

	"gorm.io/gorm"
)

// Fortune 运势分析模型
type Fortune struct {
	gorm.Model
	UserID       uint      `json:"user_id"`
	Date         time.Time `json:"date"`
	OverallScore int       `json:"overall_score"` // 总体运势指数（0-100）
	LoveScore    int       `json:"love_score"`    // 感情运势指数
	CareerScore  int       `json:"career_score"`  // 事业运势指数
	HealthScore  int       `json:"health_score"`  // 健康运势指数
	WealthScore  int       `json:"wealth_score"`  // 财运指数
	Analysis     string    `json:"analysis"`      // AI分析报告
	Suggestions  string    `json:"suggestions"`   // 建议和注意事项
}
