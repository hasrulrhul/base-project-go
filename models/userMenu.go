package models

import (
	"time"

	"gorm.io/gorm"
)

type UserMenu struct {
	ID        uint           `json:"id"`
	RoleID    uint           `gorm:"not null" json:"role_id"`
	MenuID    uint           `gorm:"not null" json:"menu_id"`
	Read      string         `gorm:"type:enum('1', '0');default:'1'" json:"read"`
	Create    string         `gorm:"type:enum('1', '0');default:'1'" json:"create"`
	Update    string         `gorm:"type:enum('1', '0');default:'1'" json:"update"`
	Delete    string         `gorm:"type:enum('1', '0');default:'1'" json:"delete"`
	Report    string         `gorm:"type:enum('1', '0');default:'1'" json:"report"`
	Role      Role           `json:"role"`
	Menu      Role           `json:"menu"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
