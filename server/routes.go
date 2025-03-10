package server

import (
	"github.com/vviveksharma/taskTracker-CLI/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	task := app.Group("/task")

	task.Get("/", h.ListTask)
	task.Post("/", h.CreateTask)
	task.Delete("/", h.DeleteTask)
	task.Put("/", h.UpdateTask)
}
