package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/thebackendgrip/grpc-hello-world/hello"
	"google.golang.org/grpc"
)

type HelloWorldServer struct {
	hello.UnimplementedHelloWordServer
}

func (s HelloWorldServer) Greet(ctx context.Context, in *hello.GreetRequest) (*hello.GreetResponse, error) {
	return &hello.GreetResponse{
		Message: fmt.Sprintf("Hello! world, %s", in.Name),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal("error occured starting server: ", err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloWordServer(s, HelloWorldServer{})
	log.Printf("gRPC server starting at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
