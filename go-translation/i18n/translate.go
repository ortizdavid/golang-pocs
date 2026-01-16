package i18n

import (
	"strings"
	"github.com/gofiber/fiber/v2"
)

func T(c *fiber.Ctx, key string, params ...map[string]string) string {
	once.Do(LoadTranslations)

	// get from header
	lang := c.Get("Accept-Language", "en")
	if len(lang) > 2 {
		lang = lang[:2]
	}

	msgs, found := translations[lang]
	if !found {
		msgs = translations["en"]
	}

	message, found := msgs[key]
	if !found {
		message, found = translations["en"][key]
		if !found {
			return key
		}
	}

	// param replacements: {name} -> value
	if len(params) > 0 {
		for k, v := range params[0] {
			message = strings.ReplaceAll(message, "{"+k+"}", v)
		}
	}

	return message
}