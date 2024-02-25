package routes

import (
	"net/http"

	albums "github.com/AthirsonSilva/music-streaming-api/cmd/server/handlers/albums"
	users "github.com/AthirsonSilva/music-streaming-api/cmd/server/handlers/users"
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
	router.Use(middlewares.RateLimiter)

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "/swagger/index.html", http.StatusFound)
	})

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	router.Route("/api/v1/albums", AlbumRoutes)
	router.Route("/api/v1/users", UserRoutes)

	return router
}

func AlbumRoutes(router chi.Router) {
	router.Use(middlewares.VerifyAuthentication)
	router.Get("/", albums.FindAllAlbums)
	router.Get("/{id}", albums.FindOneAlbumById)
	router.Post("/", albums.CreateAlbum)
	router.Put("/{id}", albums.UpdateAlbumById)
	router.Delete("/{id}", albums.DeleteAlbumById)
}

func UserRoutes(router chi.Router) {
	router.Get("/{id}", users.FindOneUserById)
	router.Post("/signup", users.SignUp)
	router.Post("/signin", users.SignIn)
	router.Get("/verify", users.VerifyUser)
}
