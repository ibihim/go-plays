package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself

	fmt.Printf(`
%s %s?%s %s
Host: %s
User-Agent: %s
Accept: %s
Accept-Encoding: %s
Accept-Charset: %s


`,
		r.Method, r.URL.Path, r.URL.RawQuery, r.Proto,
		r.Host, strings.Join(r.Header["User-Agent"], ", "),
		strings.Join(r.Header["Accept"], ", "),
		strings.Join(r.Header["Accept-Encoding"], ", "),
		strings.Join(r.Header["Accept-Charset"], ", "),
	)

	fmt.Fprintf(w, "Hello World!") // send data to client side
}

func main() {
	http.HandleFunc("/", helloWorld) // set router

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
