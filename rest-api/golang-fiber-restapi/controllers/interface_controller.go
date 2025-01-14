package controllers

import "github.com/gofiber/fiber/v2"

type InterfaceController interface {
	getAll(ctx *fiber.Ctx) error
	getTask(ctx *fiber.Ctx) error
	create(ctx *fiber.Ctx) error
	update(ctx *fiber.Ctx) error
	delete(ctx *fiber.Ctx) error
	search(ctx *fiber.Ctx) error
	RegisterRoutes(router *fiber.App)
}