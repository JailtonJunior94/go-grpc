package main

import (
	"fmt"
	"log"
	"net"

	cep "github.com/jailtonjunior94/go-grpc/pb"
	chat "github.com/jailtonjunior94/go-grpc/pb"
	"github.com/jailtonjunior94/go-grpc/services"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	cepServer := services.CepServer{}
	chatServer := services.ChatServer{}

	grpcServer := grpc.NewServer()

	cep.RegisterCepServiceServer(grpcServer, &cepServer)
	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	fmt.Println("🚀 gRPC server is running on http://localhost:9000")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over 9000: %v", err)
	}
}
