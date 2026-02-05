package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"workshop-platform/backend/internal/api/handlers"
	apimw "workshop-platform/backend/internal/api/middleware"
	"workshop-platform/backend/internal/auth"
)

type Deps struct {
	AuthHandler     *handlers.AuthHandler
	VehiclesHandler *handlers.VehiclesHandler
	JWT             *auth.JWT
}

func Router(d *Deps) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	r.Get("/health", handlers.Health)

	r.Route("/api", func(r chi.Router) {
		r.Post("/auth/login", d.AuthHandler.Login)

		r.Route("/v1", func(r chi.Router) {
			r.Use(apimw.Auth(d.JWT))
			r.Route("/vehicles", func(r chi.Router) {
				r.Get("/", d.VehiclesHandler.List)
				r.Post("/", d.VehiclesHandler.Create)
				r.Get("/{id}", d.VehiclesHandler.GetByID)
			})
		})
	})

	return r
}
