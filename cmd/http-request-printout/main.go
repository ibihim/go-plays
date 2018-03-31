package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("ReadAll Body: ", err)
		return
	}

	fmt.Printf(`
%s %s?%s %s
Host: %s
User-Agent: %s
Accept: %s
Accept-Encoding: %s
Accept-Charset: %s

Body: %s
`,
		r.Method, r.URL.Path, r.URL.RawQuery, r.Proto,
		r.Host, strings.Join(r.Header["User-Agent"], ", "),
		strings.Join(r.Header["Accept"], ", "),
		strings.Join(r.Header["Accept-Encoding"], ", "),
		strings.Join(r.Header["Accept-Charset"], ", "),
		string(body),
	)

	fmt.Fprintf(w, "Request printed to console!") // send data to client side
}

func main() {
	port := ":8080"

	http.HandleFunc("/", helloWorld) // set router

	fmt.Printf("Listening on localhost:%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
