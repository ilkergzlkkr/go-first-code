package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/illkergzlkkr/go-first-code/models"
	"github.com/illkergzlkkr/go-first-code/repositories"
)

func SetupTodoGroup(group *fiber.Router) {
	(*group).Get("/", TodosGet)
	(*group).Post("/", TodosPost)
	(*group).Get("/:id", TodoGet)
	(*group).Delete("/:id", TodoDelete)
	(*group).Patch("/:id/done", TodoMarkDone)
}

// @Tags         Todo
// @Description  get all todos
// @Produce      json
// @Param        undone query bool false "Filter Undone Todos optionally with `true`"
// @Success      200 {object} models.Todo
// @Router       /todo [get]
func TodosGet(c *fiber.Ctx) error {
	if c.QueryBool("undone") {
		return c.JSON(repositories.GetUndoneTodos())
	}

	return c.JSON(repositories.GetTodos())
}

// @Tags         Todo
// @Description  get todo with id
// @Produce      json
// @Success      200 {object} models.Todo
// @Param        id path string true "Todo Id"
// @Failure      400 {string} string
// @Failure      404 {string} string
// @Router       /todo/:id [get]
func TodoGet(c *fiber.Ctx) error {
	payload := struct {
		Id string `json:"id"`
	}{}

	if err := c.ParamsParser(&payload); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	todo := repositories.GetTodo(payload.Id)

	if todo == nil {
		return c.Status(404).SendString("Todo Not Found")
	}

	return c.JSON(todo)
}

// @Tags         Todo
// @Description  create a new todo
// @Accept       json
// @Produce      json
// @Param        title query string true "Todo Title" minlength(3)
// @Success      200 {string} string
// @Failure      400 {string} string
// @Router       /todo [post]
func TodosPost(c *fiber.Ctx) error {
	payload := struct {
		Title string `json:"title"`
	}{}

	if err := c.QueryParser(&payload); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	newTodo, err := models.NewTodo(payload.Title)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	repositories.AddTodo(newTodo)
	return c.SendString(newTodo.Id)
}

// @Tags         Todo
// @Description  delete todo with id
// @Produce      json
// @Success      204
// @Param        id path string true "Todo Id"
// @Failure      400 {string} string
// @Failure      404 {string} string
// @Router       /todo/:id [delete]
func TodoDelete(c *fiber.Ctx) error {
	payload := struct {
		Id string `json:"id"`
	}{}

	if err := c.ParamsParser(&payload); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	todo := repositories.GetTodo(payload.Id)

	if todo == nil {
		return c.Status(404).SendString("Todo Not Found")
	}

	repositories.RemoveTodo(todo.Id)
	return c.SendStatus(204)
}

// @Tags         Todo
// @Description  mark todo done
// @Produce      json
// @Success      200 {object} models.Todo
// @Param        id path string true "Todo Id"
// @Failure      400 {string} string
// @Failure      404 {string} string
// @Router       /todo/:id/done [patch]
func TodoMarkDone(c *fiber.Ctx) error {
	payload := struct {
		Id string `json:"id"`
	}{}

	if err := c.ParamsParser(&payload); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	todo := repositories.GetTodo(payload.Id)

	if todo == nil {
		return c.Status(404).SendString("Todo Not Found")
	}

	if todo.IsDone() {
		return c.Status(404).SendString("Todo is already done")
	}

	repositories.SetTodo(todo.MarkAsDone())
	return c.JSON(todo)
}
