package routes

import (
	"github.com/ezedinff/appzoid/controllers"
	"github.com/gofiber/fiber/v2"
)

func APIRoutes(app *fiber.App) {
	api := app.Group("/api")
	APIVersions(api)
}

func APIVersions(api fiber.Router) {
	v1Routes(api)
}
func v1Routes(api fiber.Router) {
	v1 := api.Group("/v1")
	userRoute(v1)
}

func userRoute(v fiber.Router) {
	// User
	user := v.Group("/user")
	user.Get("/:id", controllers.GetUser)
	user.Post("/", controllers.CreateUser)
	user.Patch("/:id", controllers.UpdateUser)
	user.Delete("/:id", controllers.DeleteUser)
}
