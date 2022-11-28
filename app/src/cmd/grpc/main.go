package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	server "github.com/seki-shinnosuke/grpc-gateway/gen/go/grpc"
	usecase "github.com/seki-shinnosuke/grpc-gateway/usecase"
)

func main() {
	srv := grpc.NewServer()
	api, err := usecase.NewTodoUsecase()
	if err != nil {
		log.Fatal("failed to new ExampleAPIServer")
	}

	server.RegisterTodoApiServer(srv, api)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen(tcp, :50051)")
	}

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("err has occured while serving: %v", err)
	}
}
