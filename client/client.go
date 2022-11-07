// Team members
// Gerson Fialho | jgfn1
// Arthur Frade | afa4
// César Silva | accs2

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	pb "github.com/afa4/golang-grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartClient(numberOfRequests int) (error) {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	serviceClient := pb.NewIsEvenServiceClient(conn)
	for i := 0; i < numberOfRequests; i++ {
		costumerId := int32(i + 1)
		fmt.Printf("Customer %d going to sushi bar\n", costumerId)
		_, err := serviceClient.IsEven(context.Background(), &pb.IsEvenRequest{Integer: costumerId})
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func main() {
	numberOfRequests, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Fatal error")
	}

	err = StartClient(numberOfRequests)
	if err != nil {
		panic("Fatal error")
	}
}
