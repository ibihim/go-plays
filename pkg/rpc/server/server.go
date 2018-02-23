package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/ibihim/go-plays/pkg/rpc/contract"
)

// StartServer starts a rpc server on 8080
func StartServer(port int) {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Unable to listen on given port: %s", err)
	}

	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

// HelloWorldHandler is a rpc handler
type HelloWorldHandler struct{}

// HelloWorld replies message with with args.name
func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello" + args.Name
	return nil
}
