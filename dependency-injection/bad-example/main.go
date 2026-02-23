package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Instantiating Infrastructure directly in Main
	logger, _ := zap.NewProduction()

	dsn := "host=localhost user=gopher password=pass dbname=pocs port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database")
	}

	cache := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Instantiating Repositories manually
	userRepo := &UserRepository{db: db}
	productRepo := &ProductRepository{db: db}

	// Instantiating Services manually
	userService := &UserService{repo: userRepo, logger: logger}
	productService := &ProductService{repo: productRepo, cache: cache}

	// Instantiating Handlers manually
	userHandler := &UserHandler{svc: userService}
	productHandler := &ProductHandler{svc: productService}

	app := fiber.New()

	// Registering Routes directly in Main
	api := app.Group("/api")

	// Routes
	userHandler.Routes(api)
	productHandler.Routes(api)

	logger.Info("Server starting on :3000")
	log.Fatal(app.Listen(":3000"))
}
