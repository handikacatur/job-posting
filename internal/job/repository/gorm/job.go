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

func (j *JobRepository) GetJobs(ctx context.Context, req request.GetJobsRequest) ([]entity.JobToCompany, error) {
	var result []entity.JobToCompany

	query := j.db.Model(&entity.Job{}).
		Select("jobs.id as job_id, companies.name as company, jobs.title as title, jobs.description as description, jobs.created_at as created_at").
		Joins("left join companies on companies.id = jobs.company_id").
		Where("(jobs.title @@ to_tsquery(?) OR jobs.description @@ to_tsquery(?))", req.Keyword, req.Keyword)

	if req.CompanyName != "" {
		query.Where("companies.name ilike ?", "%"+req.CompanyName+"%")
	}

	query.Order("jobs.created_at desc")

	if err := query.Find(&result).Error; err != nil {
		log.Errorf("error when get jobs Err: %v", err)
		return result, err
	}

	return result, nil
}

func (j *JobRepository) CreateJob(ctx context.Context, req request.CreateJobRequest) error {
	company := entity.Company{}

	if err := j.db.FirstOrCreate(&company, entity.Company{Name: req.CompanyName}).Error; err != nil {
		log.Errorf("error when first or create company")
		return err
	}

	job := entity.Job{
		CompanyID:   company.ID,
		Title:       req.Title,
		Description: req.Description,
	}
	if err := j.db.Create(&job).Error; err != nil {
		log.Errorf("error when create job")
		return err
	}

	return nil
}
