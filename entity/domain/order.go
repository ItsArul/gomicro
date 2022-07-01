package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Buyer    string `json:"buyer" form:"buyer" validate:"required"`
	Product  string `json:"product" form:"product" validate:"required"`
	Quantity int    `json:"quantity" form:"quantity" validate:"required"`
}
