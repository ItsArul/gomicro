package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	NameCategory Product   `gorm:"foreignKey:Category" json:"name" form:"name" validate:"required"`
	Product      []Product `gorm:"foreignKey:Name" json:"product" form:"product"`
}
