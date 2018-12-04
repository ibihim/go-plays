package cookie

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func cookieMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if len(req.Cookies()) == 0 {
			http.SetCookie(res, &http.Cookie{
				Name:     "diff-paths",
				Value:    uuid.NewV4().String(),
				HttpOnly: true,
				Path:     "/",
			})
		}

		f(res, req)
	}
}

func logHeader(title string, h http.Header) {
	var hStr string
	for key, value := range h {
		hStr = fmt.Sprintf("%s\n - %s: %s", hStr, key, value)
	}

	log.Printf(`


== %s =======================================
%s
`, title, hStr)
}

func loggerMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logHeader("request", req.Header)
		f(res, req)
		logHeader("response", res.Header())
	}
}

func NewDiffPathServer() {
	port := ":8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/a", loggerMiddleware(cookieMiddleware(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Response on /a")
	})))

	mux.HandleFunc("/b", loggerMiddleware(cookieMiddleware(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Response on /b")
	})))

	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
