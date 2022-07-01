package user

import "github.com/ItsArul/gomicro/entity/domain"

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	Update(id uint, user domain.User) (domain.User, error)
	Delete(id uint) error
}
