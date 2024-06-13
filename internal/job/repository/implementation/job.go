package implementation

import (
	"context"

	"github.com/handikacatur/jobs-api/internal/job/model/entity"
	"github.com/handikacatur/jobs-api/internal/job/model/request"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (j *JobRepository) GetJobs(ctx context.Context, request request.GetJobsRequest) ([]entity.Job, error) {
	return []entity.Job{}, nil
}
