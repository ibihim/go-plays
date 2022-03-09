package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if err := app("/tmp/echo.sock"); err != nil {
		log.Fatal(err)
	}
}

func app(sockAddr string) error {
	if err := os.RemoveAll(sockAddr); err != nil {
		return err
	}

	l, err := net.Listen("unix", sockAddr)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go echoServer(conn)
	}
}

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	fmt.Fprintln(c, "welcome to this unix domain socket server")

	io.Copy(c, c)
	c.Close()
}
