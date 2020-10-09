package main

import (
	"context"
	"log"

	"github.com/jailtonjunior94/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:3014", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewMathServiceClient(connection)

	request := &pb.NewSumRequest{
		Sum: &pb.Sum{
			A: 11,
			B: 12,
		},
	}

	response, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Erro no request:  %v", err)
	}
	log.Printf("O resultado é: %v", response.Result)
}
