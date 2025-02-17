package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"uniqueIndex;size:50"`
	Password  string         `gorm:"size:100"`
	Email     string         `gorm:"size:100"`
	Nickname  string         `gorm:"size:50"`
	Avatar    string         `gorm:"size:255"`
}
