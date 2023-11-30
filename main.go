package main

import (
	"chi-orders-api/application" // import application package
	"context"
	"fmt"
)

func main() {

	// create new app
	app := application.New()

	// start app
	err := app.Start(context.TODO())

	// handle errors
	if err != nil {
		fmt.Printf("error starting app: %s\n", err)
	}
}