package main

import (
	prodControll "github.com/ItsArul/gomicro/controller/product"
	prodRepo "github.com/ItsArul/gomicro/repository/product"
	"github.com/ItsArul/gomicro/routes"
	prodService "github.com/ItsArul/gomicro/services/product"

	userControll "github.com/ItsArul/gomicro/controller/user"
	userRepo "github.com/ItsArul/gomicro/repository/user"
	userService "github.com/ItsArul/gomicro/services/user"
)

func main() {
	prodRepository := prodRepo.NewProductRepository()
	prodServices := prodService.NewProductService(prodRepository)
	prodController := prodControll.NewProductController(prodServices)

	userRepository := userRepo.NewUserRepository()
	userServices := userService.NewUserServices(userRepository)
	userController := userControll.NewUserController(userServices)
	routes.Run(prodController, userController)
}
