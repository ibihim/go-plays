package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/ibihim/go-plays/pkg/rpc/contract"
)

func CreateClient(port int) *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}

	var reply contract.HelloWorldResponse
	if err := client.Call("HelloWorldHandler.HelloWorld", args, &reply); err != nil {
		log.Fatalf("error: %v\n", err)
	}

	return reply
}
