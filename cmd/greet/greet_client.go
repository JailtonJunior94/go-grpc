package main

import (
	"context"
	"log"

	"github.com/jailtonjunior94/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewGreeterClient(connection)

	request := &pb.HelloRequest{
		Name: "Meu nome é Jailton Jailton",
	}

	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("Erro no request:  %v", err)
	}
	log.Printf("O resultado é: %v", response.Message)
}
