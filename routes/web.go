package routes

import (
	"github.com/gofiber/fiber/v2"
)

// WebRoutes web routes
func WebRoutes(app *fiber.App) {
	app.Static("/", "views")
	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("../views/index.html")
	})
}
