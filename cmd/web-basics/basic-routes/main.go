package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct{}

func (a App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Foo!")
}

// FooApp handles greeting requests under the /bar route.
type FooApp struct{}

// ServeHTTP implements the http.Handler interface.
func (a FooApp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Foo!")
}

// BarApp handles greeting requests under the /bar route.
type BarApp struct{}

// ServeHTTP implements the http.Handler interface.
func (a BarApp) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Bar!")
}

func main() {
	// m is a *http.ServeMux which is the multiplexer (mux for short) or
	// router that will direct traffic within our service
	m := http.NewServeMux()

	// Register our handlers for different paths on the mux.
	m.Handle("/", App{})
	m.Handle("/foo", FooApp{})
	m.Handle("/bar", BarApp{})

	// Start the server using our mux. It also implements http.Handler
	log.Print("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", m))
}
