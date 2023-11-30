package application

import (
	"context"
	"fmt"
	"net/http"
)

// create an app struct to store dependencies
type App struct {
	router http.Handler
}

// function to construct new app
func New() *App {

	// create app and assign router function as handler
	app := &App{
		router: loadRoutes(),
	}

	return app
}

// function to start app
func (a *App) Start(ctx context.Context) error {

	// create server
	server := &http.Server{
		Addr: ":3000",
		Handler: a.router,
	}

	// listen and server
	err := server.ListenAndServe()

	// handle errors
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	// return nil if success
	return nil
}