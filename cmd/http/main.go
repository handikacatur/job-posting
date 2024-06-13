package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/handikacatur/jobs-api/cmd/config"
	"github.com/handikacatur/jobs-api/cmd/internal"
)

func main() {
	cfg := config.InitConfig()

	serv := internal.GetService(cfg)

	app := fiber.New(fiber.Config{
		AppName:   "Job posting API",
		BodyLimit: 10 * 1024 * 1024,
	})
	app.Use(recover.New())
	app.Use(cors.New())

	httpService := internal.InitService(serv)

	// Initialize Routers
	httpService.Job.SetRoute(app)

	log.Printf("fiber is running on port %s. Env: %v", cfg.API.Port, os.Getenv("GO_ENV"))

	log.Fatal(app.Listen(":" + cfg.API.Port))
}
