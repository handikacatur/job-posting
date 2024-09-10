package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/handikacatur/jobs-api/cmd/config"
)

func main() {
	cfg := config.InitConfig()

	//serv := internal.GetService(cfg)

	app := fiber.New(fiber.Config{
		AppName:   "Job posting API",
		BodyLimit: 10 * 1024 * 1024,
	})
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(fiberlogger.New(fiberlogger.Config{
		Format:        "${ip} ${status} - ${method} - ${path}\n",
		TimeFormat:    time.RFC3339,
		TimeZone:      "Asia/Jakarta",
		DisableColors: false,
	}))

	//httpService := internal.InitService()

	// Initialize Routers
	//httpService.Job.SetRoute(app)

	log.Printf("fiber is running on port %s. Env: %v", cfg.API.Port, os.Getenv("GO_ENV"))

	log.Fatal(app.Listen(":" + cfg.API.Port))
}
