package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/ibihim/go-plays/pkg/rpc-http/contract"
)

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func StartServer(port int) {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Unable to listne on given port: %s", err)
	}

	log.Printf("Server starting on port %v\n", port)

	http.Serve(l, nil)
}
