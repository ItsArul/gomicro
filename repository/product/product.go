package product

import (
	"errors"

	"github.com/ItsArul/gomicro/db"
	"github.com/ItsArul/gomicro/entity/domain"
)

type productrepository struct{}

func NewProductRepository() ProductRepository {
	return &productrepository{}
}

func (prod *productrepository) Create(p domain.Product) (domain.Product, error) {
	err := db.DB.Create(&p).Error
	if err != nil {
		return p, errors.New("cannot create product")
	}

	return p, errors.New("success create product")
}

func (prod *productrepository) FindById(id uint) (domain.Product, error) {
	var p domain.Product

	err := db.DB.First(&p, id).Error
	if err != nil {
		return p, errors.New("cannot find product")
	}

	return p, errors.New("find product")
}

func (prod *productrepository) FindAll() ([]domain.Product, error) {
	var pr []domain.Product

	err := db.DB.Find(&pr).Error
	if err != nil {
		return pr, errors.New("cannot find all product")
	}

	return pr, errors.New("find all product")
}

func (prod *productrepository) Update(id uint, p domain.Product) (domain.Product, error) {
	var pr domain.Product

	err := db.DB.Model(&p).Where("id = ?", id).Updates(&pr).Error
	if err != nil {
		return pr, errors.New("cannot update product")
	}

	return pr, errors.New("success update product")

}

func (prod *productrepository) Delete(id uint) error {
	var p domain.Product

	err := db.DB.Where("id = ?", id).Delete(&p).Error
	if err != nil {
		return errors.New("cannot deelte product")
	}

	return errors.New("success delete product")
}
