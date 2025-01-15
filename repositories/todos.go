package repositories

// Memory Repository
// Can be extended via database

// SQLite, MySQL, MongoDB

import (
	"github.com/illkergzlkkr/go-first-code/models"
)

var todos = make([]models.Todo, 0, 100)

func GetTodos() []models.Todo {
	return todos
}

func GetTodo(id string) *models.Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return &todo
		}
	}

	return nil
}

func RemoveTodo(id string) error {
	for i, todo := range todos {
		if todo.Id == id {
			// [1, 2,] 3, [4, 5]
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	return nil
}

func GetUndoneTodos() []models.Todo {
	filter := []models.Todo{}

	for _, todo := range todos {
		if !todo.Done {
			filter = append(filter, todo)
		}
	}

	return filter
}

func AddTodo(todo *models.Todo) {
	todos = append(todos, *todo)
}

func SetTodo(todo *models.Todo) error {
	err := RemoveTodo(todo.Id)

	if err != nil {
		return err
	}

	AddTodo(todo)
	return nil
}
