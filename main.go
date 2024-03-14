package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)

	type Status struct {
		name string
	}

	status := Status{
		name: "Connection was successful",
	}

	v1Router := chi.NewRouter()

	v1Router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, 200, status)
	})

	router.Mount("/v1", v1Router)

	fmt.Println("Server is running on port 3333")
	http.ListenAndServe(":3333", router)
}
