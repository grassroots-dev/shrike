package service

import (
	"context"
	"errors"
	"log"

	pb "github.com/grassroots-dev/shrike/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service is a type represneting a configured server.
type Service struct {
	DB      string
	Cache   string
	Storage string
}

// NewService returns a new service configured with all the resources.
func NewService(db string, cache string, storage string) *Service {
	return &Service{DB: "hello db", Cache: "hello cache", Storage: "hello storage"}
}

// SayHello implements helloworld.GreeterServer
func (s *Service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	log.Printf("DB is : %v", s.DB)
	log.Printf("Test updates Cache is : %v", s.Cache)
	log.Printf("Storage is : %v", s.Storage)
	err := errors.New("can't work with 42")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve id for created Account-> "+err.Error())
	}

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
