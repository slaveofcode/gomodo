package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slaveofcode/gomodo/app/api"
)

func BuildRoutes(app *fiber.App) {
	// registering /pinger for healthcheck
	app.Get("/pinger", api.PingerHandler)

	apiRoute := app.Group("/api")
	api.BuildRoutes(apiRoute)
}
