package request

type LoginValidation struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterValidation struct {
	Name     string `json:"name" binding:"required"`
	Username string `gorm:"unique" json:"username" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required,email"`
	Password string `json:"password"`
	RoleID   uint   `gorm:"not null" json:"role_id"`
}
