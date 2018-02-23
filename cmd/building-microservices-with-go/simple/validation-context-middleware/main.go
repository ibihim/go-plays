package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.Handle("/helloworld", newValidationHandler(newHelloWorldHandler()))

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return &validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var req helloWorldRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextString("name"), req.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

type validationContextString string

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return &helloWorldHandler{}
}

func (*helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextString("name")).(string)
	res := helloWorldResponse{Message: "Hello " + name}

	json.NewEncoder(rw).Encode(&res)
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message,omitempty"`
}
