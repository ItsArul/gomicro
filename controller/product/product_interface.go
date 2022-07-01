package product

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	FindId(ctx *fiber.Ctx) error
	FindCategory(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	OrderProduct(ctx *fiber.Ctx) error
	UploadProduct(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
