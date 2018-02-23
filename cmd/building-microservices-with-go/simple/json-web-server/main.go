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
	res := helloWorldResponse{"HelloWorld"}

	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		panic("Ooops")
	}
}

type helloWorldResponse struct {
	Message string `json:"message,omitempty"`
}
