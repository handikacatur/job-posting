package httpservice

import "github.com/handikacatur/jobs-api/internal/job/service"

type Handler struct {
	jobService service.JobServiceProvider
}

type HandlerConfig struct {
	JobService service.JobServiceProvider
}
