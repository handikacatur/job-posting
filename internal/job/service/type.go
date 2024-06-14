package service

import (
	"context"

	"github.com/handikacatur/jobs-api/internal/job/model/request"
	"github.com/handikacatur/jobs-api/internal/job/model/response"
	modelError "github.com/handikacatur/jobs-api/internal/model/model_error"
)

type JobServiceProvider interface {
	GetJobList(ctx context.Context, req request.GetJobsRequest) ([]response.GetJobsResponseData, modelError.ErrorIface)
	CreateJob(ctx context.Context, req request.CreateJobRequest) modelError.ErrorIface
}
