package main

import (
	"context"
	"log"
	"ms-account/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:4030"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := proto.NewGetDataClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Account(ctx, &proto.Request{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Resp)

}
