package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/loyalsfc/social-network/internal/database"

	_ "github.com/lib/pq"
)

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

	v1Router.Get("/todos", apiCfg.getTodosHandler)

	v1Router.Get("/todos/{todoId}", apiCfg.getTodoHandler)

	v1Router.Post("/todo", apiCfg.addTodoHandler)

	v1Router.Delete("/todo/{todoId}", apiCfg.deleteTodo)

	v1Router.Put("/todo/{todoId}", apiCfg.updateTodo)

	router.Mount("/v1", v1Router)

	fmt.Println("Server is running on port 3333")
	http.ListenAndServe(":3333", router)
}
