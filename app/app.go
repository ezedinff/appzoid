package app

import (
	"log"
	"os"

	"github.com/ezedinff/appzoid/config"
	"github.com/ezedinff/appzoid/database"
	"github.com/ezedinff/appzoid/models"
	"github.com/ezedinff/appzoid/routes"
	"github.com/gofiber/fiber/v2"
)

// App global variable for fiber app instance
var App *fiber.App

// Ctx global fiber context
var Ctx *fiber.Ctx

func bootApp() {
	App = fiber.New()
}
func Serve() {
	bootApp()
	config.LoadEnv()
	config.LoadAppConfig()
	routes.RegisterRoute(App)
	database.DatabaseInit()
	database.GetDB().Debug().AutoMigrate(&models.User{})
	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "3000"
	}
	err := App.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
		panic("App not starting: " + err.Error() + "on Port: " + config.AppConfig.App_Port)
	}
}
