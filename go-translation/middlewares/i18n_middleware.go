package middlewares

import "github.com/gofiber/fiber/v2"

type I18nMiddleware struct {
}

func NewI18nMiddleware() *I18nMiddleware {
	return &I18nMiddleware{}
}

func (mid *I18nMiddleware) Handle(c *fiber.Ctx) error {
	lang := c.Get("Accept-Language", "en")

	c.Locals("lang", lang)

	return c.Next()
}