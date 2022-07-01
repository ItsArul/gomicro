package product

import (
	"github.com/ItsArul/gomicro/entity/domain"
)

type ProductRepository interface {
	Create(p domain.Product) (domain.Product, error)
	FindById(id uint) (domain.Product, error)
	FindAll() ([]domain.Product, error)
	Update(id uint, p domain.Product) (domain.Product, error)
	Delete(id uint) error
}
