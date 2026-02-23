package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/services"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(svc *services.ProductService) *ProductHandler {
	return &ProductHandler{service: svc}
}

func (h *ProductHandler) Routes(router fiber.Router) {
	group := router.Group("/products")
	group.Get("/", h.getAll)
	group.Post("/", h.create)
	group.Patch("/:id", h.update)
	group.Post("/:id", h.delete)
	group.Get("/:id", h.getByID)
}