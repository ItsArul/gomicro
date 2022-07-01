package product

import "github.com/ItsArul/gomicro/entity/domain"

type ProductService interface {
	FindById(id uint) (domain.Product, error)
	FindByCategory(category string) (domain.Product, error)
	OrderProduct(order domain.Order) (domain.Order, error)
	UpdateProduct(id uint, prod domain.Product) (domain.Product, error)
	UploadProduct(prod domain.Product) (domain.Product, error)
	FindAll() ([]domain.Product, error)
	Delete(id uint) error
}
