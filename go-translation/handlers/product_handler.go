package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/go-translation/i18n"
	"github.com/ortizdavid/golang-pocs/go-translation/models"
	"github.com/ortizdavid/golang-pocs/go-translation/repositories"
)


type ProductHandler struct {
	repository *repositories.ProductRepository
}

func NewProductHandler(repository *repositories.ProductRepository) *ProductHandler {
	return &ProductHandler{
		repository: repository,
	}
}

func (h *ProductHandler) Routes(router *fiber.App) {
	group := router.Group("/products")
	group.Post("/", h.create)
	group.Put("/:id", h.update)
	group.Delete("/:id", h.delete)
	group.Get("/:id", h.getByID)
}

func (h *ProductHandler) create(c *fiber.Ctx) error {
	var product models.ProductModel
	c.BodyParser(&product)

	return c.Status(201).JSON(fiber.Map{
		"message": i18n.T(c, "product.created"),
	})
}

func (h *ProductHandler) update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product models.ProductModel
	c.BodyParser(&product)
	product.ID = id

	if err := h.repository.Update(product); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": i18n.T(c, "product.not_found"),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": i18n.T(c, "product.updated"),
	})
}

func (h *ProductHandler) delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.repository.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": i18n.T(c, "product.error.delete"),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": i18n.T(c, "product.delete"),
	})
}

func (h *ProductHandler) getByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product, err := h.repository.GetByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": i18n.T(c, "product.not_found"),
		})
	}

	return c.Status(200).JSON(product)
}