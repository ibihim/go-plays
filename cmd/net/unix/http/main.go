package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	if err := app("/tmp/echo.sock"); err != nil {
		fmt.Println(err)
	}
}

func app(sockAddr string) error {
	if err := os.RemoveAll(sockAddr); err != nil {
		return err
	}

	server := http.Server{Handler: http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "hello to the http unix socket listener")

			io.Copy(w, r.Body)
		},
	)}

	unixListener, err := net.Listen("unix", sockAddr)
	if err != nil {
		return err
	}

	return server.Serve(unixListener)
}
