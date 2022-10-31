package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/weverton-souza/go-todo-task-crud/configs"
	"github.com/weverton-souza/go-todo-task-crud/handlers"
)

func main() {
	log.Println("Iniciando")
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Get("/", handlers.GetAll)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe((fmt.Sprintf(":%s", configs.GetServerPort())), r)
}
