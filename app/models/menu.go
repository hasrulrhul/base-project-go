package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID        uint           `json:"id"`
	Parent    string         `json:"parent" binding:"required"`
	Name      string         `gorm:"not null" json:"name" binding:"required"`
	Icon      string         `json:"icon"`
	Url       string         `json:"url" binding:"required"`
	Index     uint16         `json:"index" binding:"required,number"`
	Active    string         `gorm:"type:enum('1', '0');default:'1'" json:"active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *Menu) TableName() string {
	return "menu"
}
