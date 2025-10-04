package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naufalilyasa/todolist-be-golang/configs"
	"github.com/naufalilyasa/todolist-be-golang/internal/database"
	"github.com/naufalilyasa/todolist-be-golang/internal/handlers"
	"github.com/naufalilyasa/todolist-be-golang/internal/middleware"
	"github.com/naufalilyasa/todolist-be-golang/internal/repositories"
	"github.com/naufalilyasa/todolist-be-golang/internal/services"
)

func main() {
	cfg := configs.LoadConfig()
	database.Connect(cfg)

	repoTodo := repositories.NewTodoRepository(database.DB)
	serviceTodo := services.NewTodoService(repoTodo)
	handlersTodo := handlers.NewTodoHandler(serviceTodo)

	repoCategory := repositories.NewCategoryRepository(database.DB)
	serviceCategory := services.NewCategoryService(repoCategory)
	handlersCategory := handlers.NewCategoryHandler(serviceCategory)

	r := chi.NewRouter()

	r.Use(middleware.CORS())

	r.Route("/api/todos", func(r chi.Router) {
		r.Get("/", handlersTodo.GetTodos)
		r.Get("/{id}", handlersTodo.GetTodoById)
		r.Post("/", handlersTodo.Create)
		r.Put("/{id}", handlersTodo.Update)
		r.Patch("/{id}/complete", handlersTodo.UpdateComplete)
		r.Delete("/{id}", handlersTodo.Delete)
	})

	r.Route("/api/categories", func(r chi.Router) {
		r.Get("/", handlersCategory.GetCategories)
		r.Post("/", handlersCategory.Create)
		r.Put("/{id}", handlersCategory.Update)
		r.Delete("/{id}", handlersCategory.Delete)
	})

	log.Printf("Server running at: %s", cfg.AppPort)
	http.ListenAndServe(":"+cfg.AppPort, r)
}
