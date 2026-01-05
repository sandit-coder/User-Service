package router

import (
	"UserCrud/internal/User/adapters/primary/http-adapter/handlers/user"

	"github.com/gofiber/fiber/v3"
)

func AppendUserRoutes(h *handlers.UserHandler, app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/users", h.Create)
	api.Get("/users/:id", h.GetById)
	api.Delete("/users/:id", h.DeleteById)
	api.Put("/users/:id", h.Update)
}
