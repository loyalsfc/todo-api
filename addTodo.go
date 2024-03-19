package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/loyalsfc/social-network/internal/database"
)

type parameters struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func (apiCfg apiCfg) getTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := apiCfg.DB.GetTodos(r.Context())

	if err != nil {
		errorResponse(w, 400, "An error occured")
		return
	}

	jsonResponse(w, 200, handleTodosToTodos(todos))
}

func (apiCfg apiCfg) getTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")

	id, err := uuid.Parse(todoId)

	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Error in parsing id %v", err))
	}

	requestedTodo, err := apiCfg.DB.GetTodo(r.Context(), id)

	if err != nil {
		errorResponse(w, 404, fmt.Sprintf("An Error occured: %v", err))
		return
	}

	jsonResponse(w, 200, handleTodoToTodo(requestedTodo))
}

func (apiCfg apiCfg) addTodoHandler(w http.ResponseWriter, r *http.Request) {

	decorder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decorder.Decode(&params)

	if err != nil {
		errorResponse(w, 500, "Params error")
		return
	}

	todo, err := apiCfg.DB.AddTodo(r.Context(), database.AddTodoParams{
		ID:          uuid.New(),
		Title:       params.Title,
		Description: params.Description,
		IsCompleted: params.IsDone,
	})

	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Error parsing json %v", err))
		return
	}

	jsonResponse(w, 200, handleTodoToTodo(todo))

}

func (apiCfg apiCfg) deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")

	id, err := uuid.Parse(todoId)

	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Invalid id: %v", err))
		return
	}

	apiCfg.DB.DeleteTodo(r.Context(), id)

	jsonResponse(w, 200, "Todo Deleted successfully")

}

func (apiCfg apiCfg) updateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Invalid form data: %v", err))
	}

	id, err := uuid.Parse(todoId)

	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Invalid todo id: %v", err))
		return
	}

	apiCfg.DB.UpdateTodo(r.Context(), database.UpdateTodoParams{
		ID:          id,
		Title:       params.Title,
		Description: params.Description,
		IsCompleted: params.IsDone,
	})

	jsonResponse(w, 200, "Todo updated successfully")
}
