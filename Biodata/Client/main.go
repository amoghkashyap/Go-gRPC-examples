package main

import (
	pb "github.com/amoghkashyap/Go-gRPC-examples/Biodata/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const(
	address = "localhost:17002"
)

func main() {
	conn, err := grpc.Dial(address,grpc.WithInsecure())
	if err != nil {
		log.Fatal(" error  %v", err)
	}

	defer conn.Close()

	c := pb.NewInfoClient(conn)
	res, err := c.Details(context.Background(), &pb.Biorequest{Name:"Amogh",EmailId:"amogh.kashyap@nokia.com",Age:"22"})
	if err != nil {
		log.Fatal("error  %v", err)
	}
	log.Println(res.Status)
	log.Println(res.StatusResponse)
}
