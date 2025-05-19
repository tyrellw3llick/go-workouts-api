package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tyrellw3llick/femProject/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)

		r.Get("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleGetWorkoutById))
		r.Post("/workouts", app.Middleware.RequireUser(app.WorkoutHandler.HandleCreateWorkout))
		r.Put("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleUpdateWorkoutById))
		r.Delete("/workouts/{id}", app.Middleware.RequireUser(app.WorkoutHandler.HandleDeleteWorkoutByID))
	})

	r.Get("/health", app.HealthCheck)

	r.Post("/users", app.UserHandler.HandleRegisterUser)

	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

	return r
}
