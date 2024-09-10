//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/handikacatur/jobs-api/cmd/config"
	jobRepositoryGorm "github.com/handikacatur/jobs-api/internal/job/repository/gorm"
	"github.com/handikacatur/jobs-api/internal/job/service"
)

var (
	sets = wire.NewSet(
		config.InitConfig,
		wire.FieldsOf(new(config.Config), "Database"),
		initDatabase,
	)
)

func InitService() service.JobServiceProvider {
	wire.Build(
		sets,
		jobRepositoryGorm.NewJobRepository,
		wire.Struct(new(service.JobServiceConfig), "JobRepo"),
		service.NewJobService,
	)

	return nil
}

//func InitializedDB() config.DatabaseConfig {
//	wire.Build(
//		config.InitConfig,
//		wire.FieldsOf(new(config.Config), "DBConfig"),
//		//initDatabase
//	)
//	return nil
//}
