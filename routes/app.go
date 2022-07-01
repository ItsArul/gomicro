package routes

import (
	"github.com/ItsArul/gomicro/controller/product"
	conf "github.com/ItsArul/gomicro/db"
	"github.com/ItsArul/gomicro/entity/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Run(prod product.ProductController) {
	db := conf.GetConnection()
	db.AutoMigrate(&domain.Product{}, &domain.Order{}, &domain.Category{}, &domain.User{})

	app := fiber.New()

	corsConf := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}

	app.Use(corsConf)
	api := app.Group("/api/product")
	{
		api.Get("/", prod.FindAll)
		api.Post("/", prod.UploadProduct)
		api.Put("/:id", prod.UpdateProduct)
		api.Delete("/:id", prod.Delete)
		api.Post("/order", prod.OrderProduct)
		api.Get("/:id", prod.FindId)
		api.Get("/:category", prod.FindCategory)
	}

	config := conf.GetConfig()
	app.Listen(config.RunApp.Host + config.RunApp.Port)
}
