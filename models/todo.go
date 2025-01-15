package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewTodo(title string) (*Todo, error) {
	if len(title) < 3 {
		return nil, errors.New("title length must be greater than 3")
	}

	return &Todo{
		Id:        uuid.NewString(),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}, nil
}

func (todo *Todo) IsDone() bool {
	return todo.Done
}

func (todo *Todo) MarkAsDone() *Todo {
	if !todo.IsDone() {
		todo.Done = true
	}

	return todo
}
