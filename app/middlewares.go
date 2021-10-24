package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func AssignMiddlewares(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(compress.New())
}
