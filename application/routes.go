package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// loadRoutes function to load routes
func loadRoutes() *chi.Mux {

	// create chi router
	router := chi.NewRouter()

	// add middleware
	router.Use(middleware.Logger)

	// create route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return router
}