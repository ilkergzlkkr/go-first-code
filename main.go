package main

import (
	"log"

	"github.com/illkergzlkkr/go-first-code/controllers"
	_ "github.com/illkergzlkkr/go-first-code/docs"

	"github.com/gofiber/fiber/v2"
)

//	@title        Simple Go Fiber App
//	@version      1.0.0
//	@description  Simplest API written in go

func main() {
	app := fiber.New()
	controllers.SetupSwagger(app)

	root := app.Group("/")
	controllers.SetupRootGroup(&root)

	todo := app.Group("/todo")
	controllers.SetupTodoGroup(&todo)

	log.Fatal(app.Listen(":3000"))
}
