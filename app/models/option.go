package models

import (
	"time"

	"gorm.io/gorm"
)

type Option struct {
	ID          uint           `gorm:"primary_key:auto_increment" json:"id"`
	Code        string         `json:"code" binding:"required"`
	Value       string         `json:"value" binding:"required"`
	Description string         `json:"description"`
	Index       uint16         `json:"index" binding:"required,number"`
	Active      string         `gorm:"not null;default:'1'" json:"active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
