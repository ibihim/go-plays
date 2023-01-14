package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	pb "github.com/ibihim/go-plays/cmd/net/grpc/api"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTimeServer
}

func (s *Server) GetTime(ctx context.Context, req *pb.TimeRequest) (*pb.TimeResponse, error) {
	l, err := time.LoadLocation(req.Timezone)
	if err != nil {
		return nil, err
	}

	return &pb.TimeResponse{
		Time: time.Now().In(l).String(),
	}, nil
}

func (s *Server) Listen(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	g := grpc.NewServer()
	pb.RegisterTimeServer(g, s)

	return g.Serve(l)
}

func main() {
	addr := flag.String("addr", ":50051", "the address listening to")
	s := Server{}

	flag.Parse()
	log.Println("Listening at", *addr)

	if err := s.Listen(*addr); err != nil {
		panic(err)
	}
}
