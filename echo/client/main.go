package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/k-nasa/grpc-sample/echo/proto"

	"google.golang.org/grpc"
)

func main() {
	target := "localhost:8080"
	connection, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer connection.Close()

	client := pb.NewEchoServiceClient(connection)
	msg := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := client.Echo(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		log.Println(err)
	}

	log.Println(r.GetMessage())
}
