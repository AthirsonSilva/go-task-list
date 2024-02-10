package routes

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/internal/handlers"
	"github.com/AthirsonSilva/music-streaming-api/internal/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.WriteToConsole)

	router.Route("/api/v1/albums", func(router chi.Router) {
		router.Get("/", handlers.FindAll)
		router.Get("/{id}", handlers.FindOne)
		router.Post("/", handlers.Create)
		router.Put("/{id}", handlers.Update)
		router.Delete("/{id}", handlers.Delete)
	})

	return router
}
