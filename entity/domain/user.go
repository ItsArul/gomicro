package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Password []byte `json:"password" form:"password" validate:"required,min=8"`
	Email    string `json:"email" form:"email" validate:"email"`
	Avatar   string `json:"avatar" form:"avatar"`
	Address  string `json:"address" form:"address" validate:"required"`
	Job      string `json:"job" form:"job" validate:"required"`
	Phone    int    `json:"phone" form:"phone" validate:"required"`
}
