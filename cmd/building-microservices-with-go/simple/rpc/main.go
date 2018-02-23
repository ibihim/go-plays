package main

import (
	"fmt"

	"github.com/ibihim/go-plays/pkg/rpc/client"
	"github.com/ibihim/go-plays/pkg/rpc/server"
)

func main() {
	port := 8000

	go server.StartServer(port)

	c := client.CreateClient(port)
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
