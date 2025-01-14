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

type TaskController struct {
}

func (task TaskController) RegisterRoutes(router *fiber.App) {
	jwt := security.NewAuthMiddleware(config.ApiSecret())
	group := router.Group("/api/tasks")
	group.Get("/", jwt, task.getAll)
	group.Post("/", jwt, task.create)
	group.Get("/:id", jwt, task.getTask)
	group.Put("/:id", jwt, task.update)
	group.Delete("/:id", jwt, task.delete)
	group.Get("/search/:param", jwt, task.search)
	
}

func (TaskController) getAll(ctx *fiber.Ctx) error {
	tasks := models.TaskModel{}.FindAll()
	count := len(tasks)
	if count == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": true,
			"Count": count,
			"Message": "No Tasks Found",
			"Status": "Fail",
		})
	} else {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Count": count,
			"Message": "All Tasks Found",
			"Status": "Success",
			"Data": tasks,
		})
	}
}

func (TaskController) getTask(ctx *fiber.Ctx) error {
	var taskModel models.TaskModel
	id := ctx.Params("id")
	intId := helpers.ConvertToInt(id)
	task := taskModel.FindById(intId)

	if !taskModel.Exists(intId) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": true,
			"Message": "Task Does Not Exists",
			"Status": "Fail",
		})
	} else {

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "Task Found",
			"Status": "Success",
			"Data": task,
		})
	}
}


func (TaskController) create(ctx *fiber.Ctx) error {
	var taskModel models.TaskModel
	task := new(entities.Task)
	
	err := ctx.BodyParser(task)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	} else {
		taskModel.Create(*task)
		log.Printf("Task '%s' Created ", task.TaskName)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "Task Created Successfully",
			"Status": "Success",
			"Data": task,
		})
	}
}

func (TaskController) update(ctx *fiber.Ctx) error {
	var taskModel models.TaskModel
	task := new(entities.Task)
	id := ctx.Params("id")
	task.TaskId = helpers.ConvertToInt(id)

	err := ctx.BodyParser(task)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	} else {
		taskModel.Update(*task)
		log.Printf("Task '%s' Updated", task.TaskName)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "Task Updated Successfully",
			"Status": "Success",
			"Data": task,
		})
	}
}

func (TaskController) delete(ctx *fiber.Ctx) error {
	var taskModel models.TaskModel
	id := ctx.Params("id")
	intId := helpers.ConvertToInt(id)
	task := models.TaskModel{}.FindById(intId)
	
	if !taskModel.Exists(intId) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Error": true,
			"Message": "Task Does Not Exists",
			"Status": "Fail",
		})
	} else {
		taskModel.Delete(intId)
		log.Printf("Task '%s' Deleted ", task.TaskName)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"Error": false,
			"Message": "Task Deleted",
			"Status": "Success",
		})
	}
}

func (TaskController) search(ctx *fiber.Ctx) error {
	param := ctx.Params("param")
	results := models.TaskModel{}.Search(param)
	count := len(results)

	log.Printf("Search for Task '%v' and %v Results Founds", param, count)
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
