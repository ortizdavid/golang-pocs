package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/ortizdavid/golang-fiber-restapi/config"
	"github.com/ortizdavid/golang-fiber-restapi/models"
	"github.com/ortizdavid/golang-fiber-restapi/security"
)

type AuthController struct {
}


func (auth AuthController) RegisterRoutes(router *fiber.App) {
	jwt := security.NewAuthMiddleware(config.ApiSecret())
	group := router.Group("/api/auth")
	group.Post("/login", auth.login)
	group.Get("/logout", auth.logout)
	group.Get("/protected", jwt, auth.protected)
}

func (AuthController) login(ctx *fiber.Ctx) error {
	var userModel models.UserModel
	loginRequest := new(security.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := userModel.FindByCredentials(loginRequest.UserName, loginRequest.Password)
	if !userModel.Exists(user.UserId) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User Unauthorized",
		})
	}

	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"UserId":    user.UserId,
		"UserName": user.UserName,
		"Exp":   time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.ApiSecret()))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(security.LoginResponse {
		Token: t,
	})
}

func (AuthController) logout(ctx *fiber.Ctx) error {
	return ctx.SendString("About API")
}


// Protected route
func (AuthController) protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	userName := claims["UserName"].(string)
	return c.SendString("Welcome  " + userName)
}