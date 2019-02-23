package main

import (
	"fmt"
	"log"
	"net/http"
)

// Default is the default greeting.
func Default(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}

// Foo greets requests at the /foo route.
func Foo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Foo!")
}

// Bar greets requests at the /bar route.
func Bar(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello Bar!")
}

func main() {
	// Use http.HandleFunc instead of http.Handle. Instead of requiring a
	// full-blown type implementing http.Handler this just wants any function
	// that accepts a response writer and request.
	http.HandleFunc("/", Default)
	http.HandleFunc("/foo", Foo)
	http.HandleFunc("/bar", Bar)

	log.Print("Listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
