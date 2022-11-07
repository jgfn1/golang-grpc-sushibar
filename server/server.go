// Team members
// Gerson Fialho | jgfn1
// Arthur Frade | afa4
// César Silva | accs2

package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/afa4/golang-grpc/protos"
	"google.golang.org/grpc"
)

type isEvenService struct {
	pb.UnimplementedIsEvenServiceServer
}

func (*isEvenService) IsEven(context context.Context, req *pb.IsEvenRequest) (*pb.IsEvenReply, error) {
	// fmt.Printf("IsEven remote call for input %d\n", req.Integer)
	fmt.Printf("Customer %d arrived and waiting\n", req.Integer)
	if(req.Integer % int32(5) == int32(0)){
		fmt.Printf("\nFriends %d, %d, %d, %d, %d are now eating sushi\n\n", req.Integer, 
		req.Integer - 1, req.Integer - 2, req.Integer - 3 , req.Integer - 4)
		return &pb.IsEvenReply{IsEven: req.Integer%2 == 0}, nil
	}
	fmt.Printf("Customer %d sitting\n", req.Integer)
	return &pb.IsEvenReply{IsEven: req.Integer%2 == 0}, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func StartServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterIsEvenServiceServer(grpcServer, &isEvenService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	StartServer()
}
