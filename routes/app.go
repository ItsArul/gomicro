package routes

import (
	"github.com/ItsArul/gomicro/controller/product"
	"github.com/ItsArul/gomicro/controller/user"
	conf "github.com/ItsArul/gomicro/db"
	"github.com/ItsArul/gomicro/entity/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Run(prod product.ProductController, user user.UserController) {
	db := conf.GetConnection()
	db.AutoMigrate(&domain.Product{}, &domain.Order{}, &domain.Category{}, &domain.User{})

	app := fiber.New()

	corsConf := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}

	app.Use(corsConf)
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}",
	}))

	app.Get("/monitoring", monitor.New(monitor.Config{Title: "Monitor My App"}))

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

	apiUsr := app.Group("/api/user")
	{
		apiUsr.Post("/register", user.Register)
		apiUsr.Post("/login", user.Login)
		apiUsr.Put("/:id", user.Update)
		apiUsr.Delete("/:id", user.Delete)
	}

	config := conf.GetConfig()
	app.Listen(config.RunApp.Host + config.RunApp.Port)
}
