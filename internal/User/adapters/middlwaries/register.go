package middleware

import "github.com/gofiber/fiber/v3"

func Register(app *fiber.App, slogMw fiber.Handler) {
	app.Use(slogMw)
}
git