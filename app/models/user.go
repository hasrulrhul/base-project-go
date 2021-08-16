package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Username  string         `gorm:"unique" json:"username"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"password"`
	RoleID    uint           `gorm:"not null" json:"role_id"`
	Role      Role           `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b *User) TableName() string {
	return "user"
}
