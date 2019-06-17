package main

import (
	pb "github.com/k-nasa/grpc-sample/echo/proto"
	"google.golang.org/grpc"

	"log"
	"net"
)

func main() {
	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, &echoService{})

	log.Printf("start server on port %s\n", port)
	if err := server.Serve(lis); err != nil {
		log.Printf("failed to serve!")
	}
}
