package controllers

import "github.com/gofiber/fiber/v2"

var Version = "1.0.0"

func SetupRootGroup(group *fiber.Router) {
	(*group).Get("/", RootController)
	(*group).Get("/version", VersionController)
}

// @Description  get the welcome message
// @Tags         Base
// @Success      200 {string} string
// @Router       / [get]
func RootController(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Simplest API written in go!")
}

// @Description  get the version
// @Tags         Base
// @Success      200 {string} string
// @Router       /version [get]
func VersionController(c *fiber.Ctx) error {
	return c.SendString(Version)
}
