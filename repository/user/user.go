package user

import (
	"errors"

	"github.com/ItsArul/gomicro/db"
	"github.com/ItsArul/gomicro/entity/domain"
)

type userrepository struct{}

func NewUserRepository() UserRepository {
	return &userrepository{}
}

func (usr *userrepository) Create(user domain.User) (domain.User, error) {
	err := db.DB.Create(&user).Error
	if err != nil {
		return user, errors.New("cannot create user")
	}

	return user, errors.New("success create user")
}

func (usr *userrepository) Update(id uint, user domain.User) (domain.User, error) {
	var usrs domain.User
	err := db.DB.Model(&user).Where("id = ?", id).Updates(&usrs)
	if err != nil {
		return usrs, errors.New("cannot update user")
	}

	return usrs, errors.New("success update user")
}

func (usr *userrepository) Delete(id uint) error {
	var user domain.User
	err := db.DB.Where("id = ?", id).Delete(&user)
	if err != nil {
		return errors.New("cant delete user")
	}

	return errors.New("success delete user")
}
