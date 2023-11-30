package main

import (
	"net/http"
)

func main() {

	// create server
	server := &http.Server{
		Addr: ":3000", // port
		Handler: http.HandlerFunc(basicHandler), // interface when server recieves request
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