package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/handikacatur/jobs-api/cmd/config"
)

func main() {
	var cmd string
	flag.StringVar(&cmd, "cmd", "migrate", "give command up/rollback/down/seed")
	flag.Parse()

	appConfig := config.InitConfig()

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		appConfig.Database.Driver,
		appConfig.Database.Username,
		appConfig.Database.Password,
		appConfig.Database.Host,
		appConfig.Database.Port,
		appConfig.Database.Name,
	)

	db, err := sql.Open(appConfig.Database.Driver, dsn)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://tools/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	switch cmd {
	case "up":
		// migrate database
		log.Println("Starting process migration...")

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to run migrate up: %v", err)
		}
	case "down":
		// running command down migration db. -- only works on development
		log.Println("Starting process downing database...")

		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to run migrate down: %v", err)
		}
	}

	log.Println("Migrations applied successfully")
}
