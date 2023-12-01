package application

import (
	"chi-orders-api/model"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database connection attempts
const CONN_ATTEMPTS = 10

// create an app struct to store dependencies
type App struct {
	router http.Handler
	psql *gorm.DB
	config Config
}

// function to construct new app
func New(config Config) *App {

	var connStr string

	// check environment
	if os.Getenv("APP_ENV") == "production" {
		
		// set production (docker) database
		PG_PASS := os.Getenv("POSTGRES_PASSWORD")
		PG_USER := os.Getenv("POSTGRES_USER")
		PG_DB := os.Getenv("POSTGRES_DB")

		// connection string
		connStr = fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable", PG_USER, PG_PASS, PG_DB)

	} else {
		connStr = "host=localhost dbname=chi-orders-db port=5432 sslmode=disable"
	}

	var psql *gorm.DB
	var dbConnected bool = false

	// attempt to connect to database multiple times
	for i := 0; i < CONN_ATTEMPTS && !dbConnected; i++ {

		// create psql connection
		var err error
		psql, err = gorm.Open(postgres.Open(connStr))

		// check psql errors
		if err != nil {
			fmt.Printf("[-] psql connection error: %s\n", err.Error())
			time.Sleep(10 * time.Second)
		} else {
			fmt.Println("[+] psql connected")
			dbConnected = true
		}
	}

	// check if database connected
	if !dbConnected {
		panic("[-] failed to connect to database")
	}

	// migrate models
	err := psql.AutoMigrate(&model.Order{}, &model.LineItem{})
	if err != nil {
		panic(err)
	}

	// create app and assign router function as handler
	app := &App{
		psql: psql,
		config: config,
	}

	// load routes
	app.loadRoutes()

	return app
}

// function to start app
func (a *App) Start(ctx context.Context) error {

	// create server
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", a.config.ServerPort),
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
