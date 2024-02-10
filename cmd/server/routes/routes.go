package routes

import (
	"net/http"

	albums "github.com/AthirsonSilva/music-streaming-api/cmd/server/handlers/albums"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/middlewares"
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
		router.Get("/", albums.FindAll)
		router.Get("/{id}", albums.FindOne)
		router.Post("/", albums.Create)
		router.Put("/{id}", albums.Update)
		router.Delete("/{id}", albums.Delete)
	})

	return router
}
