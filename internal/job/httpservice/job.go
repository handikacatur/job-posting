package httpservice

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/handikacatur/jobs-api/internal/job/model/request"
	"github.com/handikacatur/jobs-api/internal/job/model/response"
	errSvc "github.com/handikacatur/jobs-api/internal/model/response"
)

func NewHandler(cfg HandlerConfig) *Handler {
	return &Handler{
		jobService: cfg.JobService,
	}
}

func (h *Handler) GetJobList(c *fiber.Ctx) error {
	resp, err := h.jobService.GetJobList(c.Context(), request.GetJobsRequest{})
	if err != nil {
		return c.Status(err.GetHttpCode()).JSON(errSvc.Error{
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
