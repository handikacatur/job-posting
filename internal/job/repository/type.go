package repository

import (
	"context"

	"github.com/handikacatur/jobs-api/internal/job/model/entity"
	"github.com/handikacatur/jobs-api/internal/job/model/request"
)

type JobRepositoryProvider interface {
	GetJobs(ctx context.Context, request request.GetJobsRequest) ([]entity.Job, error)
}
