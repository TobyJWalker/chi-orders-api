package main

import (
	"net/http" // http package

	"github.com/go-chi/chi/v5"            // chi plugin
	"github.com/go-chi/chi/v5/middleware" // chi middleware
)

func main() {

	// create chi router
	router := chi.NewRouter()

	// add middleware
	router.Use(middleware.Logger) // log requests

	// create route
	router.Get("/", basicHandler)

	// create server
	server := &http.Server{
		Addr: ":3000", // port
		Handler: router, // interface when server recieves request
	}

	// start server
	err := server.ListenAndServe()

	// check for errors
	if err != nil {
		panic(err)
	}

}

// create basic http handler
func basicHandler(w http.ResponseWriter, r *http.Request) { // w = response writer, r = request
	w.Write([]byte("Hello World")) // write to response writer
}