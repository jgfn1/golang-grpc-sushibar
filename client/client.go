package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

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

	var source = rand.NewSource(time.Now().UnixNano())
	var rand = rand.New(source)
	integer := rand.Int31()
	reply, err := serviceClient.IsEven(context.Background(), &pb.IsEvenRequest{Integer: integer})
	if err != nil {
		log.Fatalln("Failed to call IsEven service")
	}
	fmt.Printf("IsEven remote response for input %d = %t\n", integer, reply.GetIsEven())
}

func main() {
	StartClient()
}
