package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/slaveofcode/gomodo/app"
	"github.com/slaveofcode/gomodo/app/helpers/logger"
)

func init() {
	logger.Init(log.DebugLevel)
}

func main() {
	fiberApp := fiber.New(fiber.Config{
		ServerHeader: "Gomodo App", // TODO: Replace with Your own project name
	})

	app.AssignMiddlewares(fiberApp)
	app.BuildRoutes(fiberApp)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := fiberApp.Listen(":" + port)

	if err != nil {
		log.Fatalln("Unable to Start API:", err.Error())
	}
}
