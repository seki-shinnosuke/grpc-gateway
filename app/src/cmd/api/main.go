package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	server "github.com/seki-shinnosuke/grpc-gateway/gen/go/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcServerAddress = "localhost:50051"

func main() {
	grpcGateway := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := server.RegisterTodoApiHandlerFromEndpoint(context.Background(), grpcGateway, grpcServerAddress, opts); err != nil {
		log.Fatal("failed to register grpc-server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcGateway)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("err")
	}
}
