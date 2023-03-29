package main

import (
	"database/sql"
	"net"

	"github.com/jailtonjunior94/go-grpc/internal/database"
	"github.com/jailtonjunior94/go-grpc/internal/pb"
	"github.com/jailtonjunior94/go-grpc/internal/service"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
