package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// index Godoc
// @Summary Home page of API
// @Description Starting Page
// @Tags Home
// @Produce plain
// @Success 200 {string} string
// @Router / [get]
func Index(c *fiber.Ctx) error {
	return c.SendString("Hello. Welcome to gofiber app!")
}

// GetAPIRoutes Godoc
// @Summary Get API Routes
// @Description Gets list of routes
// @Tags Api
// @Produce json
// @Success 200 {object} serializers.Response
// @Router /api/ [get]
func GetAPIRoutes(c *fiber.Ctx) error {
	data, _ := json.MarshalIndent(c.App().GetRoutes(true), "", "  ")
	return c.JSON(string(data))
}
