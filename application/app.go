package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	} else {
		fmt.Println("[+] psql connected")
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

	// create channel to listen for errors from server goroutine
	ch := make(chan error, 1) // 1 is buffer size

	// goroutine to start server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("error starting server: %w", err) // send error to channel
		}

		// close the channel if error
		close(ch)
	}()

	// select statement to listen for context cancellation or error from server goroutine
	select {
	
	// graceful shutdown with timeout
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return server.Shutdown(timeout)
	
	// error from server goroutine
	case err := <-ch:
		return err
	}
}
