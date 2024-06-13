package internal

import (
	"os"

	"github.com/handikacatur/jobs-api/cmd/config"

	jobRepository "github.com/handikacatur/jobs-api/internal/job/repository/gorm"
	"github.com/handikacatur/jobs-api/internal/job/service"
)

func GetService(config config.Config) *Service {
	db := initDatabase(config.Database)

	jobRepo := jobRepository.NewJobRepository(db)

	jobService := service.NewJobService(service.JobServiceConfig{
		JobRepo: jobRepo,
	})

	serv := &Service{
		Env: os.Getenv("GO_ENV"),
		Job: jobService,
	}

	return serv
}
