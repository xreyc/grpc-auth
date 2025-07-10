package server

import (
	"log"
	"net"

	authv1 "github.com/xreyc/grpc-auth/internal/gen/go/auth/v1"
	userHandler "github.com/xreyc/grpc-auth/internal/handler/grpc"
	"google.golang.org/grpc"
)

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authv1.RegisterUserServiceServer(grpcServer, userHandler.NewUserHandler())

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
