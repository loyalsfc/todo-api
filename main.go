package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type Todos struct {
	Id     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func main() {
	todos := []Todos{
		{
			Id:     uuid.New().String(),
			Task:   "See a woman",
			IsDone: false,
		},
		{
			Id:     uuid.New().String(),
			Task:   "Read a book",
			IsDone: false,
		},
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	type Status struct {
		Name string
	}

	status := Status{
		Name: "Connection was successful",
	}

	v1Router := chi.NewRouter()

	v1Router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, 200, status)
	})

	v1Router.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, 200, todos)
	})

	v1Router.Get("/todos/{todoId}", func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")

		requestedTodo := []Todos{}

		for _, todo := range todos {
			if todo.Id == todoId {
				requestedTodo = append(requestedTodo, todo)
				break
			}
		}

		if len(requestedTodo) > 0 {
			jsonResponse(w, 200, requestedTodo[0])
		} else {
			errorResponse(w, 404, "Todo not found")
		}

	})

	v1Router.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Task   string `json:"task"`
			IsDone bool   `json:"is_done"`
		}

		decorder := json.NewDecoder(r.Body)

		params := parameters{}

		err := decorder.Decode(&params)

		if err != nil {
			errorResponse(w, 500, "Params error")
			return
		}

		todos = append(todos, Todos{
			Id:     uuid.New().String(),
			Task:   params.Task,
			IsDone: params.IsDone,
		})

		jsonResponse(w, 200, todos)

	})

	v1Router.Delete("/todo/{todoId}", func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")

		temporalTodo := []Todos{}

		for _, todo := range todos {
			if todoId != todo.Id {
				temporalTodo = append(temporalTodo, todo)
			}
		}

		todos = temporalTodo
		fmt.Println(temporalTodo)
		jsonResponse(w, 200, "Todo Deleted successfully")

	})

	v1Router.Put("/todo/{todoId}", func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")

		type Param struct {
			Task   string `json:"task"`
			IsDone bool   `json:"is_done"`
		}

		decoder := json.NewDecoder(r.Body)

		params := Param{}

		err := decoder.Decode(&params)

		if err != nil {
			errorResponse(w, 400, "Invalid form data")
		}

		for index, todo := range todos {
			if todoId == todo.Id {
				todos[index] = Todos{
					Id:     todoId,
					Task:   params.Task,
					IsDone: params.IsDone,
				}
			}
		}

		jsonResponse(w, 200, "Todo updated successfully")
	})

	router.Mount("/v1", v1Router)

	fmt.Println("Server is running on port 3333")
	http.ListenAndServe(":3333", router)
}
