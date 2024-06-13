package service

import (
	"context"

	"github.com/handikacatur/jobs-api/internal/job/model/request"
	"github.com/handikacatur/jobs-api/internal/job/model/response"
	"github.com/handikacatur/jobs-api/internal/job/repository"
	modelError "github.com/handikacatur/jobs-api/internal/model/model_error"
)

type JobServiceConfig struct {
	JobRepo repository.JobRepositoryProvider
}

type jobService struct {
	jobRepo repository.JobRepositoryProvider
}

func NewJobService(cfg JobServiceConfig) JobServiceProvider {
	return &jobService{jobRepo: cfg.JobRepo}
}

func (j *jobService) GetJobList(ctx context.Context, req request.GetJobsRequest) ([]response.GetJobsResponseData, modelError.ErrorIface) {
	var resp []response.GetJobsResponseData

	jobs, err := j.jobRepo.GetJobs(ctx, req)
	if err != nil {
		return resp, modelError.New(modelError.ErrorCodeInternalServer)
	}

	for _, job := range jobs {
		resp = append(resp, response.GetJobsResponseData{
			CompanyName: "Company", // TODO: THIS SHIT
			Title:       job.Title,
			Description: job.Description,
		})
	}

	return resp, nil
}
