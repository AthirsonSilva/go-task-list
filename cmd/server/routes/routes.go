package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	task "github.com/AthirsonSilva/music-streaming-api/cmd/server/handlers/tasks"
	users "github.com/AthirsonSilva/music-streaming-api/cmd/server/handlers/users"
	"github.com/AthirsonSilva/music-streaming-api/cmd/server/middlewares"
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

	router.Route("/api/v1/task", TaskRoutes)
	router.Route("/api/v1/users", UserRoutes)

	return router
}

func TaskRoutes(router chi.Router) {
	router.Use(middlewares.VerifyAuthentication)
	router.Get("/", task.FindAllTasks)
	router.Get("/{id}", task.FindOneTaskById)
	router.Post("/", task.CreateTask)
	router.Put("/{id}", task.UpdateTaskById)
	router.Delete("/{id}", task.DeleteTaskById)
}

func UserRoutes(router chi.Router) {
	router.Get("/{id}", users.FindOneUserById)
	router.Post("/signup", users.SignUp)
	router.Post("/signin", users.SignIn)
	router.Get("/verify", users.VerifyUser)
}
