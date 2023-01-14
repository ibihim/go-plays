package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/ibihim/go-plays/cmd/net/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	tz := flag.String("tz", "UTC", "timezone of current time")
	flag.Parse()

	cnn, err := grpc.Dial(
		*addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cnn.Close()

	c := pb.NewTimeClient(cnn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.GetTime(ctx, &pb.TimeRequest{
		Timezone: *tz,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Time)
}
