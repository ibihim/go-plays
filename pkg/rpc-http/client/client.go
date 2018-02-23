package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/ibihim/go-plays/pkg/rpc-http/contract"
)

func CreateClient(port int) *rpc.Client {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatalf("dialing: ", err)
	}
	return client
}

func PerformRequest(c *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse
	if err := c.Call("HelloWorldHandler.HelloWorld", args, &reply); err != nil {
		log.Fatalf("error: %v", err)
	}

	return reply
}
