package services

import (
	"log"

	protos "github.com/jailtonjunior94/go-grpc/pb"
	"golang.org/x/net/context"
)

// Server - Estrutura de conexão gRPC
type ChatServer struct{}

// SayHello - Método de Hello World com gRPC
func (s *ChatServer) SayHello(ctx context.Context, message *protos.Message) (*protos.Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &protos.Message{Body: "Hello From the Server!"}, nil
}
