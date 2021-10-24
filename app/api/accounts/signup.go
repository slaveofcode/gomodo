package accounts

import "github.com/gofiber/fiber/v2"

func SignUpHandler(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"ok": true,
		},
	)

	return nil
}
