package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{service: svc}
}

func (h *UserHandler) Routes(router fiber.Router) {
	group := router.Group("/users")
	group.Get("/", h.getAll)
	group.Post("/", h.create)
	group.Patch("/:id", h.update)
	group.Post("/:id", h.delete)
	group.Get("/:id", h.getByID)
}