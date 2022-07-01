package product

import (
	"errors"

	"github.com/ItsArul/gomicro/db"
	"github.com/ItsArul/gomicro/entity/domain"
	"github.com/ItsArul/gomicro/repository/product"
)

type productservice struct {
	p product.ProductRepository
}

func NewProductService(prod product.ProductRepository) ProductService {
	return &productservice{
		p: prod,
	}
}

func (pr *productservice) FindById(id uint) (domain.Product, error) {
	prod, err := pr.p.FindById(id)
	if err != nil {
		return prod, errors.New("cannot find product")
	}

	return prod, errors.New("success find product")
}

func (pr *productservice) FindByCategory(category string) (domain.Product, error) {

	err := db.DB.Where("category = ?", category).Error
	if err != nil {
		return domain.Product{}, errors.New("cannot find category")
	}

	return domain.Product{}, errors.New("cannot find category")
}

func (pr *productservice) OrderProduct(order domain.Order) (domain.Order, error) {
	var product domain.Product
	var user domain.User

	err := db.DB.Model(&product).Where("name = ?", order.Product).Error
	if err != nil {
		return order, errors.New("cannot find product")
	} else {
		orderProd := domain.Order{
			Buyer:    user.Name,
			Product:  product.Name,
			Quantity: order.Quantity,
		}

		db.DB.Create(&orderProd)
		product.Quantity = product.Quantity - order.Quantity
		db.DB.Where("name = ?", order.Product).Updates(&product)
	}

	return order, errors.New("success order")
}

func (pr *productservice) UpdateProduct(id uint, prod domain.Product) (domain.Product, error) {
	p, err := pr.p.Update(id, prod)
	if err != nil {
		return p, errors.New("cannot update product")
	}

	return p, errors.New("seuccess update product")
}

func (pr *productservice) UploadProduct(prod domain.Product) (domain.Product, error) {
	p, err := pr.p.Create(prod)
	if err != nil {
		return p, errors.New("cannot upload product")
	}

	return p, errors.New("success upload product")
}

func (pr *productservice) FindAll() ([]domain.Product, error) {

	var product domain.Product
	p, err := pr.p.FindAll()
	if err != nil {
		return p, errors.New("cannot find all product")
	} else {
		category := domain.Category{
			NameCategory: product,
			Product:      p,
		}

		db.DB.Create(&category)
	}

	return p, errors.New("success find all")
}

func (pr *productservice) Delete(id uint) error {
	err := pr.p.Delete(id)
	if err != nil {
		return errors.New("cannot delete product")
	}

	return errors.New("SUCCESS delete product")
}
