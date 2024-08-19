package main

import (
	"context"
	"log"
	"time"

	"github.com/thebackendgrip/grpc-hello-world/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not create a grpc connection", err)
	}
	defer conn.Close()
	c := hello.NewHelloWordClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	greetResponse, err := c.Greet(ctx, &hello.GreetRequest{
		Name: "thebackendgrip",
	})
	if err != nil {
		log.Fatal("could not send gRPC greet request: ", err)
	}

	log.Printf("response from gRPC server: %s", greetResponse.Message)
}
