package main

import (
	"go-project/api/grpc/server"
	"go-project/pkg/repositories"
	"go-project/pkg/services"
	pb "go-project/proto_gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("grpc server started on port: 50051")

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userServer := server.NewUserGrpcServer(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
