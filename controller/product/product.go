package product

import (
	"strconv"

	"github.com/ItsArul/gomicro/entity/domain"
	"github.com/ItsArul/gomicro/entity/web"
	"github.com/ItsArul/gomicro/services/product"
	"github.com/gofiber/fiber/v2"
)

type productcontroller struct {
	p product.ProductService
}

func NewProductController(prod product.ProductService) ProductController {
	return &productcontroller{
		p: prod,
	}
}

func (pr *productcontroller) FindId(ctx *fiber.Ctx) error {
	getId := ctx.Params("id")
	id, _ := strconv.Atoi(getId)

	prod, err := pr.p.FindById(uint(id))
	if err != nil {
		response := web.FailFindId{
			Id:      uint(id),
			Message: "cannot find product",
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	responses := web.SuccessFindId{
		NameProduct: prod.Name,
		Description: prod.Description,
		Price:       prod.Price,
	}

	return ctx.Status(fiber.StatusOK).JSON(responses)
}

func (pr *productcontroller) FindCategory(ctx *fiber.Ctx) error {
	getCategory := ctx.Params("category")

	prod, err := pr.p.FindByCategory(getCategory)
	if err != nil {
		responses := web.FailFindCategory{
			Category: getCategory,
			Message:  "cannot find category",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(responses)
	}

	response := web.SuccessFindCategory{
		Category: getCategory,
		Product:  prod.Name,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pr *productcontroller) UpdateProduct(ctx *fiber.Ctx) error {
	getId := ctx.Params("id")
	id, _ := strconv.Atoi(getId)
	var product domain.Product

	errProd := ctx.BodyParser(product)
	if errProd != nil {
		fiber.NewError(fiber.StatusBadRequest, "cannot parser")
	}

	prod, err := pr.p.UpdateProduct(uint(id), product)
	if err != nil {
		response := web.FailUpdate{
			Id:      uint(id),
			Message: "cannot update product",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := web.SuccessUpdate{
		Id:      uint(id),
		Product: prod,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pr *productcontroller) OrderProduct(ctx *fiber.Ctx) error {
	var order domain.Order
	ord, err := pr.p.OrderProduct(order)
	if err != nil {
		response := web.FailOrder{
			Id:      ord.ID,
			Message: "cannot order product",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := fiber.Map{
		"data": ord,
	}

	ctx.Bind(response)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pr *productcontroller) UploadProduct(ctx *fiber.Ctx) error {
	var prod domain.Product

	p, err := pr.p.UploadProduct(prod)
	if err != nil {
		response := web.FailUpload{
			Id:      p.ID,
			Name:    p.Name,
			Message: "cannot update product",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := web.SuccessUpload{
		NameProduct: p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
	}

	ctx.Bind(fiber.Map{"data": p})

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pr *productcontroller) FindAll(ctx *fiber.Ctx) error {

	prod, err := pr.p.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("cant find all product")
	}

	response := fiber.Map{
		"data": prod,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (pr *productcontroller) Delete(ctx *fiber.Ctx) error {
	getId := ctx.Params("id")
	id, _ := strconv.Atoi(getId)

	err := pr.p.Delete(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("CANT DELETE PRODUCT")
	}
	return ctx.Status(fiber.StatusOK).JSON("success deleter product")
}
