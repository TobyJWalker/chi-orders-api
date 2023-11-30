package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"chi-orders-api/handler"
)

// loadRoutes function to load routes
func loadRoutes() *chi.Mux {

	// create chi router
	router := chi.NewRouter()

	// add middleware
	router.Use(middleware.Logger)

	// create index route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// orders route
	router.Route("/orders", loadOrderRoutes)

	return router
}

// loadOrderRoutes function to load order routes
func loadOrderRoutes(router chi.Router) {

	// create order handler
	orderHandler := &handler.Order{}

	// create routes for http methods
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.Delete)

}