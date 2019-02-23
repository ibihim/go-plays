package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()

	upgrader := websocket.Upgrader{}
	http.HandleFunc("/echo", makeEcho(upgrader))
	http.HandleFunc("/", handleHome)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func makeEcho(upgrader websocket.Upgrader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cnn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}
		defer cnn.Close()

		for {
			msgType, msg, err := cnn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("recv: %s\n", msg)

			err = cnn.WriteMessage(msgType, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
