package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Owner       string `json:"owner" form:"owner" validate:"required"`
	Avatar      string `json:"avatar" form:"avatar" validate:"required"`
	Quantity    int    `json:"quantity" form:"quantity" validate:"required"`
	Category    string `json:"category" form:"category" validate:"required"`
}
