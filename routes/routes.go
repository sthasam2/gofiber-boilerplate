package routes

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"app/controllers"
	_ "app/docs"
	routes "app/routes/api"
)

func SetupRoutes(app *fiber.App) {

	// Routes

	app.Get("/", controllers.Index)

	// api/
	api := app.Group("/api")
	api.Get("/", controllers.GetAPIRoutes).Name("API entry")
	api.Get("/docs/swagger/*", fiberSwagger.WrapHandler).Name("Swagger Docs")

	//
	// Forwarding
	routes.Path_auth(api.Group("/auth"))

}
