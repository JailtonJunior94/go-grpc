package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/jailtonjunior94/go-grpc/chat"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer connection.Close()

	chatClient := chat.NewChatServiceClient(connection)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := chatClient.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
