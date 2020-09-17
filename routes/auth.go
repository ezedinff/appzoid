package routes

import (
	"github.com/ezedinff/appzoid/controllers"
	"github.com/gofiber/fiber/v2"
)

// AuthRoutes auth routes
func AuthRoutes(app *fiber.App) {
	app.Post("/login", controllers.Login)
}
