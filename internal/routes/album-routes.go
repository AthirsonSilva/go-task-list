package routes

import (
	"net/http"

	"github.com/AthirsonSilva/music-streaming-api/internal/handlers"
	"github.com/AthirsonSilva/music-streaming-api/internal/middlewares"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.WriteToConsole)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	router.Route("/api/v1/albums", func(router chi.Router) {
		router.Get("/", handlers.FindAll)
		router.Get("/{id}", handlers.FindOne)
		router.Post("/", handlers.Create)
		router.Put("/{id}", handlers.Update)
		router.Delete("/{id}", handlers.Delete)
	})

	return router
}
