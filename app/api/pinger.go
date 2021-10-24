package api

import "github.com/gofiber/fiber/v2"

func PingerHandler(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"ok": true,
		},
	)

	return nil
}
