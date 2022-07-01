package user

import (
	"strconv"
	"time"

	"github.com/ItsArul/gomicro/db"
	"github.com/ItsArul/gomicro/entity/domain"
	userServ "github.com/ItsArul/gomicro/services/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type usercontroller struct {
	user userServ.UserService
}

func NewUserController(user userServ.UserService) UserController {
	return &usercontroller{
		user: user,
	}
}

func (usr *usercontroller) Login(ctx *fiber.Ctx) error {
	var user domain.User
	config := db.GetConfig()

	err := ctx.BodyParser(user)
	if err != nil {
		return fiber.ErrBadRequest
	}

	hashes := []byte(user.Password)

	us, errs := usr.user.Login(user, hashes)
	if errs != nil {
		response := fiber.Map{
			"Id":      us.ID,
			"Message": "cannot login",
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	claims := jwt.MapClaims{
		"Name":  user.Name,
		"Email": user.Email,
		"Exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	T, errs := token.SignedString([]byte(config.JWTSecret))
	if errs != nil {
		return ctx.Status(fiber.StatusForbidden).JSON("jwt error")
	}

	response := fiber.Map{
		"Name":  us.Name,
		"Email": us.Email,
		"Token": T,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (us *usercontroller) Register(ctx *fiber.Ctx) error {
	var user domain.User

	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("cant register")
	}

	usr, errs := us.user.Register(user)
	if errs != nil {
		response := fiber.Map{
			"Id":      usr.ID,
			"Message": "cannot register",
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := fiber.Map{
		"Name":  usr.Name,
		"Email": usr.Email,
		"Job":   usr.Job,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (us *usercontroller) Update(ctx *fiber.Ctx) error {
	var user domain.User
	getId := ctx.Params("id")
	id, _ := strconv.Atoi(getId)

	errUser := ctx.BodyParser(user)
	if errUser != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON("cannot binding user")
	}

	usr, err := us.user.Update(uint(id), user)
	if err != nil {
		response := fiber.Map{
			"Id":      usr.ID,
			"Message": "cannot update user",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := fiber.Map{
		"Name":    usr.Name,
		"Email":   usr.Email,
		"Message": "success update user",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (us *usercontroller) Delete(ctx *fiber.Ctx) error {
	getId := ctx.Params("id")
	id, _ := strconv.Atoi(getId)

	err := us.user.Delete(uint(id))
	if err != nil {
		response := fiber.Map{
			"Message": "cannot delete user",
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := fiber.Map{
		"Message": "success delete user",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
