package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"chi-orders-api/handler"
	"chi-orders-api/repository/order"
)

// loadRoutes function to load routes
func (a *App) loadRoutes() {

	// create chi router
	router := chi.NewRouter()

	// add middleware
	router.Use(middleware.Logger)

	// create index route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// orders route
	router.Route("/orders", a.loadOrderRoutes)

	a.router = router
}

// loadOrderRoutes function to load order routes
func (a *App) loadOrderRoutes(router chi.Router) {

	// create order handler
	orderHandler := &handler.Order{
		Repo: &order.PostgresRepo{
			Client: a.psql,
		},
	}

	// create routes for http methods
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.Delete)

}