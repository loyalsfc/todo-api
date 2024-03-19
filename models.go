package main

import (
	"github.com/google/uuid"
	"github.com/loyalsfc/social-network/internal/database"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   string    `json:"created_at"`
}

func handleTodoToTodo(dbTodo database.Todo) (todo Todo) {
	return Todo{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: dbTodo.Description,
		IsCompleted: dbTodo.IsCompleted,
		CreatedAt:   dbTodo.CreatedAt.Time.String(),
	}
}

func handleTodosToTodos(dbTodos []database.Todo) (todo []Todo) {
	convertedTodo := []Todo{}

	for _, todo := range dbTodos {
		convertedTodo = append(convertedTodo, handleTodoToTodo(todo))
	}

	return convertedTodo
}
