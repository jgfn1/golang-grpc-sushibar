// Team members
// Gerson Fialho | jgfn1
// Arthur Frade | afa4
// César Silva | accs2

package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	pb "github.com/afa4/golang-grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartClient(numberOfRequests int) (int, error) {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	serviceClient := pb.NewIsEvenServiceClient(conn)

	var rttsSum = 0
	for i := 0; i < numberOfRequests; i++ {
		var source = rand.NewSource(time.Now().UnixNano())
		var rand = rand.New(source)
		integer := rand.Int31n(9)
		start := time.Now().UnixNano()
		_, err := serviceClient.IsEven(context.Background(), &pb.IsEvenRequest{Integer: integer})
		end := time.Now().UnixNano()
		if err != nil {
			panic(err)
		}

		rtt := end - start
		rttsSum += int(rtt)
	}
	var rttMean = (rttsSum / numberOfRequests)
	return rttMean, nil
}

func main() {
	numberOfRequests, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Fatal error")
	}

	rttMean, err := StartClient(numberOfRequests)
	if err != nil {
		panic("Fatal error")
	}

	fmt.Println(rttMean)
}
