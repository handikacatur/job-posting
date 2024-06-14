package httpservice

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/handikacatur/jobs-api/internal/job/model/request"
	"github.com/handikacatur/jobs-api/internal/job/model/response"
	errModel "github.com/handikacatur/jobs-api/internal/model/response"
)

func NewHandler(cfg HandlerConfig) *Handler {
	return &Handler{
		jobService: cfg.JobService,
	}
}

func (h *Handler) GetJobList(c *fiber.Ctx) error {
	query := new(request.GetJobsRequest)
	if err := c.QueryParser(query); err != nil {
		return c.Status(http.StatusBadRequest).JSON(errModel.ResponseStatusOnly{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	resp, err := h.jobService.GetJobList(c.Context(), *query)
	if err != nil {
		return c.Status(err.GetHttpCode()).JSON(errModel.Error{
			StatusCode: err.GetHttpCode(),
			Message:    err.GetErrorCodeMessage().Error(),
			ErrorCode:  err.GetErrorCode(),
		})
	}

	return c.Status(http.StatusOK).JSON(response.GetJobsResponse{
		StatusCode: http.StatusOK,
		Message:    "Get job list success!",
		Data:       resp,
	})
}
