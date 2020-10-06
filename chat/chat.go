package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server - Estrutura de conexão gRPC
type Server struct{}

// SayHello - Método de Hello World com gRPC
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
