package main

import (
	"fmt"
	"log"
	"net/http"
)

// App provides application level context for our handler
type App struct{}

// ServeHTTP implements the http.Handler interface. It gives the same
// greeting to every request. ResponseWriter is an interface and does
// not need to be referenced as a pointer. Request is a struct and
// therefor needs a pointer reference.
func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}

func main() {
	// Create a value of our App. It implements http.Handler so we can
	// pass it to http.ListenAndServe
	var a App

	// Log that the server is start5ing so we see it's alive.
	log.Print("Listening on :3000. Ctr-c to cancel.")

	// Start the http server to handle the request.
	log.Fatal(http.ListenAndServe(":3000", a))
}
