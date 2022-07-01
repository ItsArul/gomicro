package main

import (
	prodControll "github.com/ItsArul/gomicro/controller/product"
	prodRepo "github.com/ItsArul/gomicro/repository/product"
	"github.com/ItsArul/gomicro/routes"
	prodService "github.com/ItsArul/gomicro/services/product"
)

func main() {
	prodRepository := prodRepo.NewProductRepository()
	prodServices := prodService.NewProductService(prodRepository)
	prodController := prodControll.NewProductController(prodServices)

	routes.Run(prodController)
}
