package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	var req helloWorldRequest
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res := helloWorldResponse{"Hello, " + req.Name}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message,omitempty"`
}
