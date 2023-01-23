package routes

import (
	"github.com/gofiber/fiber/v2"

	"app/controllers"
)

func Path_auth(api_router fiber.Router) {
	api_router.Post("/register", controllers.RegisterUser)
	api_router.Post("/login", controllers.LoginController)
}
