package gorm

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
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

func (j *JobRepository) GetJobs(ctx context.Context, request request.GetJobsRequest) ([]entity.JobToCompany, error) {
	var result []entity.JobToCompany

	query := j.db.Model(&entity.Job{}).
		Select("jobs.id as job_id, companies.name as company, jobs.title as title, jobs.description as description, jobs.created_at as created_at").
		Joins("left join companies on companies.id = jobs.company_id").
		Where("(jobs.title @@ to_tsquery(?) OR jobs.description @@ to_tsquery(?))", request.Keyword, request.Keyword)

	if request.CompanyName != "" {
		query.Where("companies.name ilike ?", "%"+request.CompanyName+"%")
	}
	if err := query.Find(&result).Error; err != nil {
		log.Errorf("error when get jobs Err: %v", err)
		return result, err
	}

	return result, nil
}
