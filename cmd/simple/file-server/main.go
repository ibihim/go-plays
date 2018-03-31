package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	gopher := http.FileServer(http.Dir("./tmp/images"))
	// without the StripPrefix, the server would look for imgaes in ./tmp/images/gopher
	http.Handle("/gopher/", http.StripPrefix("/gopher/", gopher))

	http.HandleFunc("/helloworld", health)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
func health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "alive\n")
}
