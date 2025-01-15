// setup swagger
package controllers

import (
	"fmt"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupSwagger(app *fiber.App) {
	app.Static("/docs/doc.json", "docs/swagger.json")
	app.Static("/docs/doc.yaml", "docs/swagger.yaml")
	app.Get("/docs/*", swagger.HandlerDefault)

	app.Get("/reference", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			Theme:   "",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		return c.Type("html").SendString(htmlContent)
	})
}
