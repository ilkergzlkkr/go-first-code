package main

import (
	"context"
	"log"

	"github.com/illkergzlkkr/go-first-code/controllers"
	_ "github.com/illkergzlkkr/go-first-code/docs"
	"github.com/illkergzlkkr/go-first-code/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

//	@title        Simple Go Fiber App
//	@version      1.0.0
//	@description  Simplest API written in go

var fiberLambda *fiberadapter.FiberLambda

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, request)
}

func main() {
	app := fiber.New()
	controllers.SetupSwagger(app)

	root := app.Group("/")
	controllers.SetupRootGroup(&root)

	todo := app.Group("/todo")
	controllers.SetupTodoGroup(&todo)

	if helpers.IsLambda() {
		fiberLambda = fiberadapter.New(app)
		lambda.Start(Handler)
	} else {
		log.Fatal(app.Listen(":3000"))
	}
}
