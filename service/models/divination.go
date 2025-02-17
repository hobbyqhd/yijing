package models

import (
	"time"

	"gorm.io/gorm"
)

type DivinationType string

const (
	TypeZodiac DivinationType = "zodiac"
	TypeTarot  DivinationType = "tarot"
	TypeYijing DivinationType = "yijing"
	TypeBazi   DivinationType = "bazi"
)

type Divination struct {
	ID         uint           `gorm:"primaryKey;column:id;autoIncrement"`
	CreatedAt  time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at"`
	UserID     uint           `gorm:"index;column:user_id;not null"`
	Type       DivinationType `gorm:"type:varchar(20);column:type;not null"`
	Question   string         `gorm:"type:text;column:question;not null"`
	Input      string         `gorm:"type:text;column:input;not null"`
	Result     string         `gorm:"type:text;column:result;not null"`
	AIAnalysis string         `gorm:"type:text;column:ai_analysis"`
}

// IsValidDivinationType 验证占卜类型是否有效
func IsValidDivinationType(t string) bool {
	switch DivinationType(t) {
	case TypeZodiac, TypeTarot, TypeYijing, TypeBazi:
		return true
	default:
		return false
	}
}
