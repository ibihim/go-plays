package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	dataAvailable := false

	http.HandleFunc("/wait", func(w http.ResponseWriter, r *http.Request) {
		log.Println("wait called")

		pusher, ok := w.(http.Pusher)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for {
			select {}
		}

		if err := pusher.Push("/data", nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		log.Println("data called")

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok {
			// Push is supported. Try pushing rather than
			// waiting for the browser request these static assets.
			if err := pusher.Push("/static/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/static/style.css", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, indexHTML)
	})

	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil))
}

const indexHTML = `<html>
<head>
	<title>Hello World</title>
	<script src="/static/app.js"></script>
	<link rel="stylesheet" href="/static/style.css"">
</head>
<body>
<script>
var source = new EventSource('/');

source.onmessage = function(e) {
  document.body.innerHTML += "SSE notification: " + e.data + '<br />';

  // fetch resource via XHR... from cache!
  var xhr = new XMLHttpRequest();
  xhr.open('GET', e.data);
  xhr.onload = function() {
	document.body.innerHTML += "Message: " + this.response + '<br />';
  };

  xhr.send();
};
</script>
</body>
</html>
`
