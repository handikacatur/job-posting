package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/handikacatur/jobs-api/cmd/config"
	"github.com/handikacatur/jobs-api/internal/job/httpservice"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
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
