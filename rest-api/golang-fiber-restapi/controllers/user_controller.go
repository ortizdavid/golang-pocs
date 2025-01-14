package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-fiber-restapi/config"
	"github.com/ortizdavid/golang-fiber-restapi/entities"
	"github.com/ortizdavid/golang-fiber-restapi/helpers"
	"github.com/ortizdavid/golang-fiber-restapi/models"
	"github.com/ortizdavid/golang-fiber-restapi/security"
)

type UserController struct {
}

func (user UserController) RegisterRoutes(router *fiber.App) {
	jwt := security.NewAuthMiddleware(config.ApiSecret())
	group := router.Group("/api/users")
	group.Get("/", jwt, user.getAll)
	group.Post("/", jwt, user.create)
	group.Get("/:id", jwt, user.getUser)
	group.Put("/:id", jwt, user.update)
	group.Delete("/:id", jwt, user.delete)
	group.Get("/search/:param", jwt, user.search)
}

func (UserController) getAll(ctx *fiber.Ctx) error {
	users := models.UserModel{}.FindAll()
	count := len(users)
	if count == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": true,
			"Count": count,
			"Message": "Users Not Found",
			"Status": "Fail",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Count": count,
			"Message": "All Users Found",
			"Status": "Success",
			"Data": users,
		})
	}
}

func (UserController) getUser(ctx *fiber.Ctx) error {
	var userModel models.UserModel
	id := ctx.Params("id")
	intId := helpers.ConvertToInt(id)
	user := userModel.FindById(intId)

	if !userModel.Exists(intId) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": true,
			"Message": "User Does Not Exists",
			"Status": "Fail",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "User Found",
			"Status": "Success",
			"Data": user,
		})
	}
}


func (UserController) create(ctx *fiber.Ctx) error {
	var userModel models.UserModel
	user := new(entities.User)
	
	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	} else {
		userModel.Create(*user)
		log.Printf("User '%s' Created ", user.UserName)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "User Created Successfully",
			"Status": "Success",
			"Data": user,
		})
	}
}

func (UserController) update(ctx *fiber.Ctx) error {
	var userModel models.UserModel
	user := new(entities.User)
	id := ctx.Params("id")
	user.UserId = helpers.ConvertToInt(id)

	err := ctx.BodyParser(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	} else {
		userModel.Update(*user)
		log.Printf("User '%s' Updated", user.UserName)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "User Updated Successfully",
			"Status": "Success",
			"Data": user,
		})
	}
}

func (UserController) delete(ctx *fiber.Ctx) error {
	var userModel models.UserModel
	id := ctx.Params("id")
	intId := helpers.ConvertToInt(id)
	user := models.UserModel{}.FindById(intId)
	
	if !userModel.Exists(intId) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": true,
			"Message": "User Does Not Exists",
			"Status": "Fail",
		})
	} else {
		userModel.Delete(intId)
		log.Printf("User '%s' Deleted ", user.UserName)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "User Deleted",
			"Status": "Success",
		})
	}
}

func (UserController) search(ctx *fiber.Ctx) error {
	param := ctx.Params("param")
	results := models.UserModel{}.Search(param)
	count := len(results)

	log.Printf("Search for User '%v' and %v Results Founds", param, count)
	if count == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": true,
			"Message": "Results Not Found",
			"Status": "Fail",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "Results Found",
			"Status": "Success",
			"Data": results,
		})
	}
}
