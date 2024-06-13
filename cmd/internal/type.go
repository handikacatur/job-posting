package internal

import (
	job "github.com/handikacatur/jobs-api/internal/job/httpservice"
	"github.com/handikacatur/jobs-api/internal/job/service"
)

type Service struct {
	Env string
	Job service.JobServiceProvider
}

type HTTPService struct {
	Job *job.Handler
}
