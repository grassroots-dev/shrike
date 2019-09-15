package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/grassroots-dev/shrike/api"
	"github.com/grassroots-dev/shrike/service"

	"github.com/grassroots-dev/shrike/interceptors"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string
	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string
	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9091", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "8080", "HTTP port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "localhost", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "tern", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "tern", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "tern", "Database schema")
	flag.Parse()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptors.LoggingInterceptor,
	),
	),
	)
	pb.RegisterGreeterServer(s, service.NewService("db", "cache", "storage"))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
