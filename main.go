package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/loyalsfc/social-network/internal/database"

	_ "github.com/lib/pq"
)

type Todos struct {
	Id     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type apiCfg struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	postgressKey := os.Getenv("GOOSE_TOKEN")

	if postgressKey == "" {
		log.Fatal("Connection string is empty")
	}

	conn, err := sql.Open("postgres", postgressKey)

	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	db := database.New(conn)

	apiCfg := apiCfg{
		DB: db,
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
		todoList, err := apiCfg.DB.GetTodos(r.Context())

		if err != nil {
			errorResponse(w, 400, "An error occured")
			return
		}

		jsonResponse(w, 200, todoList)
	})

	v1Router.Get("/todos/{todoId}", func(w http.ResponseWriter, r *http.Request) {
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

		jsonResponse(w, 200, requestedTodo)

	})

	v1Router.Post("/todo", func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Task        string `json:"task"`
			Description string `json:"description"`
			IsDone      bool   `json:"is_done"`
		}

		decorder := json.NewDecoder(r.Body)

		params := parameters{}

		err := decorder.Decode(&params)

		if err != nil {
			errorResponse(w, 500, "Params error")
			return
		}

		todo, err := apiCfg.DB.AddTodo(r.Context(), database.AddTodoParams{
			ID:          uuid.New(),
			Title:       params.Task,
			Description: params.Description,
			IsCompleted: params.IsDone,
		})

		if err != nil {
			errorResponse(w, 400, fmt.Sprintf("Error parsing json %v", err))
			return
		}

		jsonResponse(w, 200, todo)

	})

	v1Router.Delete("/todo/{todoId}", func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")

		id, err := uuid.Parse(todoId)

		if err != nil {
			errorResponse(w, 400, fmt.Sprintf("Invalid id: %v", err))
			return
		}

		apiCfg.DB.DeleteTodo(r.Context(), id)

		jsonResponse(w, 200, "Todo Deleted successfully")

	})

	v1Router.Put("/todo/{todoId}", func(w http.ResponseWriter, r *http.Request) {
		todoId := chi.URLParam(r, "todoId")

		type parameters struct {
			Task        string `json:"task"`
			Description string `json:"description"`
			IsDone      bool   `json:"is_done"`
		}

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
			Title:       params.Task,
			Description: params.Description,
			IsCompleted: params.IsDone,
		})

		jsonResponse(w, 200, "Todo updated successfully")
	})

	router.Mount("/v1", v1Router)

	fmt.Println("Server is running on port 3333")
	http.ListenAndServe(":3333", router)
}
