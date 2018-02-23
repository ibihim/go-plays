package main

import (
	"fmt"

	"github.com/ibihim/go-plays/pkg/rpc-http/client"
	"github.com/ibihim/go-plays/pkg/rpc-http/server"
)

func main() {
	port := 8000

	server.StartServer(port)

	c := client.CreateClient(port)
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Printf("%v\n", reply.Message)
}
