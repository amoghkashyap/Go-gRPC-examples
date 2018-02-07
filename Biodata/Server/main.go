package main

import (
	"Go-gRPC-examples/CassandraDBMigrations/cassandra"
	"Go-gRPC-examples/CassandraDBMigrations/cassandra/entities"
	pb "github.com/amoghkashyap/Go-gRPC-examples/Biodata/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":17002"
)

var (
	isBioPresent bool
)

type server struct{}

func (s *server) Details(ctx context.Context, in *pb.Biorequest) (*pb.Bioresponse, error) {

	person := entities.Biodata{}
	person.SetName(in.GetName())
	person.SetAge(in.GetAge())
	person.SetEmailID(in.GetEmailId())

	isBioPresent := cassandra.FindBioWithName(person)

	if isBioPresent {
		return &pb.Bioresponse{StatusResponse: "Biodata already present", Status: false}, nil
	} else {
		cassandra.InsertBio(person)
		return &pb.Bioresponse{StatusResponse: "Biodata Entry Successful", Status: true}, nil
	}

}

func main() {
	log.Println("starting Biodata gRPC service")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterInfoServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal("error %v", err)
	}
}
