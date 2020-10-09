package servers

import (
	"context"
	"log"

	"github.com/jailtonjunior94/go-grpc/pb"
)

// Math -
type Math struct {
}

// Sum -
func (m *Math) Sum(ctx context.Context, req *pb.NewSumRequest) (*pb.NewSumResponse, error) {
	log.Printf("Executando na Sum: %v", req)
	res := req.Sum.A + req.Sum.B

	return &pb.NewSumResponse{Result: res}, nil
}
