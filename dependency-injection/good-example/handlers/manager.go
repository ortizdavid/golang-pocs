package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/container"
)

type HandlerManager struct {
	UserHandler    *UserHandler
	ProductHandler *ProductHandler
}

func NewHandlerManager(svcs *container.ServiceContainer) *HandlerManager {
	return &HandlerManager{
		UserHandler:    NewUserHandler(svcs.UserService),
		ProductHandler: NewProductHandler(svcs.ProductService),
	}
}

// InitRoutes centraliza todo o roteamento da API
func (m *HandlerManager) InitRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	
	m.UserHandler.Routes(api)
	m.ProductHandler.Routes(api)
}