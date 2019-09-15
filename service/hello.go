package service

import (
	"context"
	"log"
	pb "github.com/grassroots-dev/shrike/api"

)

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	log.Printf("DB is : %v", s.DB)
	log.Printf("Cache is : %v", s.Cache)
	log.Printf("Storage is : %v", s.Storage)

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
