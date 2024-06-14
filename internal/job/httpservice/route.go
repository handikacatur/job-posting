package httpservice

import "github.com/gofiber/fiber/v2"

func (h *Handler) SetRoute(app *fiber.App) {
	jobGroup := app.Group("/api/v1/jobs")
	jobGroup.Get("/", h.GetJobList)
	jobGroup.Post("/", h.CreateJob)
}
