package middleware

import (
	"github.com/ezedinff/appzoid/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() func(*fiber.Ctx) {
	jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.AppConfig.SECRET),
		ErrorHandler: jwtError,
	})
	return nil
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	} else {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
	return nil
}
