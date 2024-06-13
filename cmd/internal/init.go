package internal

import (
	"fmt"
	"log"

	"github.com/handikacatur/jobs-api/cmd/config"
	"github.com/handikacatur/jobs-api/internal/job/httpservice"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase(config config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Driver,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("database connected successfully")

	return db
}

func InitService(serv *Service) HTTPService {
	return HTTPService{
		Job: httpservice.NewHandler(httpservice.HandlerConfig{
			JobService: serv.Job,
		}),
	}
}
