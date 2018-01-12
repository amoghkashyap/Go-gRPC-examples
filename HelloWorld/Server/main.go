package main

import (
	"log"
	"net"

	pb "github.com/amoghkashyap/Go-gRPC-examples/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":17001"

type server struct{}

//This method attaches Hello to the string sent by client and returns it back
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

//This method attaches Bye to the string sent by client and returns it back
func (s *server) SayBye(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Bye " + in.GetName()}, nil
}

func main() {

	log.Println("starting gRPC service")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal("error %v", err)
	}

}
