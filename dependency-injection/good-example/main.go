package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/container"
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/handlers"
	"github.com/ortizdavid/golang-pocs/dependency-injection/good-example/resources"
)

func main() {
    //  Infra
    infra := resources.NewInfraResources()

    // Containers (DI)
    repoContainer := container.NewRepositoryContainer(infra.DB)
    svcContainer := container.NewServiceContainer(repoContainer, infra)

    // Handlers & Routes
    app := fiber.New()
    handlerManager := handlers.NewHandlerManager(svcContainer)
    handlerManager.InitRoutes(app) 

    // Start
    infra.Logger.Info("Server started on :3000")
    app.Listen(":3000")
}