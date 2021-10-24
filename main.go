package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	log "github.com/sirupsen/logrus"
	"github.com/slaveofcode/gomodo/app"
	"github.com/slaveofcode/gomodo/app/helpers/logger"
)

const (
	WEB_STATIC_DIR = ".webstatic"
)

func init() {
	logger.Init(log.DebugLevel)
}

func main() {
	fiberApp := fiber.New(fiber.Config{
		ServerHeader: "Gomodo App", // TODO: Replace with Your own project name
	})

	// Serving static files when enabled and directory exist
	serveStatic := os.Getenv("SERVE_WEBSTATIC")
	if _, err := os.Stat(WEB_STATIC_DIR); serveStatic != "" && !os.IsNotExist(err) {
		fiberApp.Static("/", WEB_STATIC_DIR)
	}

	app.AssignMiddlewares(fiberApp)
	app.BuildRoutes(fiberApp)

	// auto-proxy index request to vite webserver
	webServerAddr := os.Getenv("WEBSERVER_PORT")
	if webServerAddr != "" {
		fiberApp.Use("/", proxy.Balancer(proxy.Config{
			Servers: []string{webServerAddr},
		}))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := fiberApp.Listen(":" + port)

	if err != nil {
		log.Fatalln("Unable to Start API:", err.Error())
	}
}
