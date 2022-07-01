package user

import (
	"errors"

	"github.com/ItsArul/gomicro/db"
	"github.com/ItsArul/gomicro/entity/domain"
	"github.com/ItsArul/gomicro/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type userservice struct {
	usr user.UserRepository
}

func NewUserServices(user user.UserRepository) UserService {
	return &userservice{
		usr: user,
	}
}

func (us *userservice) Login(user domain.User, hash []byte) (domain.User, error) {

	err := db.DB.Model(&user).Where("email = ?", user.Email)
	if err != nil {
		return user, errors.New("cannot find email")
	}

	if user.ID == 0 {
		return user, errors.New("USER NOT FOUND")
	}

	errHash := bcrypt.CompareHashAndPassword(user.Password, hash)
	if errHash != nil {
		return user, errors.New("password invalid")
	}

	return user, errors.New("SUCCESS LOGIN")
}

func (us *userservice) Register(user domain.User) (domain.User, error) {
	usr, err := us.usr.Create(user)
	if err != nil {
		return usr, errors.New("cannot register")
	}

	return usr, errors.New("success register")
}

func (us *userservice) Update(id uint, user domain.User) (domain.User, error) {
	usr, err := us.usr.Update(id, user)
	if err != nil {
		return usr, errors.New("cannot update user")
	}
	return usr, errors.New("success update user")
}

func (us *userservice) Delete(id uint) error {
	err := us.usr.Delete(id)
	if err != nil {
		return errors.New("cannot delete user")
	}
	return errors.New("success delete user")
}
