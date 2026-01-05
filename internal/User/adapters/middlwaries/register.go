package middleware

import "github.com/gofiber/fiber/v3"

func Register(app *fiber.App, handlers ...fiber.Handler) {
	for _, m := range handlers {
		app.Use(m)
	}
}
