package main

import (
	"chi-orders-api/application" // import application package
	"os"

	"context"
	"fmt"
	"os/signal"
)

func main() {

	// create new app
	app := application.New(application.LoadConfig())

	// create root context with signal which listens for SIGINT
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// start app
	err := app.Start(ctx)

	// handle errors
	if err != nil {
		fmt.Printf("error starting app: %s\n", err)
	}

	// cancel context
	cancel()
}