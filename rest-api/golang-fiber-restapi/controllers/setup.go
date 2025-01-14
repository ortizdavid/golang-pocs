package controllers

import "github.com/gofiber/fiber/v2"

func SetupRoutes(router *fiber.App) {
	AuthController{}.RegisterRoutes(router)
	FrontController{}.RegisterRoutes(router)
	UserController{}.RegisterRoutes(router)
	TaskController{}.RegisterRoutes(router)
}