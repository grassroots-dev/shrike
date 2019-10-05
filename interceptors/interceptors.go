package interceptors

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

// LoggingInterceptor logs requests to standard output.
func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// Calls the handler
	h, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}
	// Logging with grpclog (grpclog.LoggerV2)
	fmt.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

// AuthInterceptor logs requests to standard output.
func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	fmt.Printf("Running authentications for rpc handler..\n")
	// Calls the handler
	h, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}
	return h, err
}

// WebHookInterceptor logs requests to standard output.
func WebHookInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Calls the handler
	h, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Triggering webhooks attached to rpc handler..\n")

	return h, err
}
