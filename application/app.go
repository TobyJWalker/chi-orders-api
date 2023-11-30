package application

import (
	"context"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// create an app struct to store dependencies
type App struct {
	router http.Handler
	psql *gorm.DB
}

// function to construct new app
func New() *App {

	// create psql connection
	psql, err := gorm.Open(postgres.Open("host=localhost dbname=chi-orders-db port=5432 sslmode=disable"))

	// check psql errors
	if err != nil {
		panic(err)
	}

	// create app and assign router function as handler
	app := &App{
		router: loadRoutes(),
		psql: psql,
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