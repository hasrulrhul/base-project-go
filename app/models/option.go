package models

import (
	"time"

	"gorm.io/gorm"
)

type Option struct {
	ID          uint           `json:"id"`
	Code        string         `json:"code"`
	Value       string         `json:"value"`
	Description string         `json:"description"`
	Index       uint16         `json:"index"`
	Active      string         `gorm:"type:enum('1', '0');default:'1'" json:"active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *Option) TableName() string {
	return "option"
}
