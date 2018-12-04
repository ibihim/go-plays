package cookie

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World!")
}

func redirect(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "http://localhost:8080", http.StatusFound)
}

func setCookie(cookieName, cookieValue string, f http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		http.SetCookie(res, &http.Cookie{
			Name:     cookieName,
			Value:    cookieValue,
			HttpOnly: true,
		})

		f(res, req)
	}
}

func Exec() {
	http.HandleFunc("/", setCookie("root", "hiewr ist 8080?", helloWorld))

	go func() {
		port := "8081"
		mux := http.NewServeMux()
		mux.HandleFunc("/", helloWorld)
		mux.HandleFunc("/8080", setCookie("808i1", "redirect inc", redirect))
		log.Printf("Listening on :%s\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
	}()

	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", setCookie("8080", "hello world!", helloWorld))
	log.Printf("Listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
