package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/afa4/golang-grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartClient() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to start client %v", err)
	}
	defer conn.Close()
	serviceClient := pb.NewIsEvenServiceClient(conn)
	reply, err := serviceClient.IsEven(context.Background(), &pb.IsEvenRequest{Integer: 122})
	if err != nil {
		log.Fatalln("Failed to call IsEven service")
	}
	fmt.Println(reply.GetIsEven())
}

func main() {
	StartClient()
}
