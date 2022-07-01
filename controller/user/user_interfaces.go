package user

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
