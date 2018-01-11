package main

import (
	pb "github.com/amoghkashyap/Go-gRPC-examples/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:17001"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(" error maga %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	res, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})

	r, er := c.SayBye(context.Background(), &pb.HelloRequest{Name: "world"})

	if err != nil || er != nil {
		log.Fatal("error  %v", err)
	}

	log.Println(res.Message)
	log.Println(r.Message)
}
