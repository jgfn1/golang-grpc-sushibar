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
