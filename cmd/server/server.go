package main

import (
	"log"
	"net"

	"github.com/jailtonjunior94/go-grpc/pb"
	"github.com/jailtonjunior94/go-grpc/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterMathServiceServer(grpcServer, &servers.Math{})
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Cannot start the gRPC server: %v", err)
	}
	log.Print("Server is running on: 3000")
	grpcServer.Serve(listener)
}
