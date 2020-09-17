package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoute(app *fiber.App) {
	APIRoutes(app)
	AuthRoutes(app)
}
