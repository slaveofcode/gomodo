package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/slaveofcode/gomodo/app/api/accounts"
)

func buildAccountRoutes(router fiber.Router) {
	router.Post("/sign-up", accounts.SignUpHandler)
	router.Post("/activate", accounts.ActivateHandler)
	router.Post("/sign-in", accounts.SignInHandler)
	router.Post("/reset-password", accounts.ResetPasswordHandler)
}

func BuildRoutes(router fiber.Router) {
	// registering routes /api/accounts
	buildAccountRoutes(router.Group("/accounts"))
}
