package user

import "github.com/ItsArul/gomicro/entity/domain"

type UserService interface {
	Login(user domain.User, hash []byte) (domain.User, error)
	Register(user domain.User) (domain.User, error)
	Update(id uint, user domain.User) (domain.User, error)
	Delete(id uint) error
}

/**
	claims := jwt.MapClaims{
		"Name":  user.Name,
		"Email": user.Email,
		"Exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	T, errs := token.SignedString([]byte(config.JWTSecret))
	if errs != nil {
		return user, errors.New("cannot claim tokens")
	}
**/
