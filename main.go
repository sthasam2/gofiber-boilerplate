package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"app/configs"
	"app/db"
	"app/models"
	"app/routes"
)

// @title           API
// @version         0.1.0
// @description     This is a gofiber sample

// @BasePath  /api

// @securityDefinitions.basic  JWT
func main() {

	app := fiber.New(fiber.Config{
		AppName: "App v.0.1.0",
	})

	/////////////////////
	// 		Configs
	/////////////////////

	config := configs.SetupConfigs()

	/////////////////////
	// 		Routes
	/////////////////////

	routes.SetupRoutes(app)

	/////////////////////
	// 		Database
	/////////////////////

	db.SetupDatabase()

	// Migrations
	db.PgDB.AutoMigrate(&models.User{})

	/////////////////////
	// 		Middlewares
	/////////////////////

	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)

}
