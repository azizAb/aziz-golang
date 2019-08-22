package main

import (
	"context"
	"ms-account/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4030")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterGetDataServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Account(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	name := request.GetName()

	result := "Hallo " + name
	return &proto.Response{Resp: result}, nil
}
