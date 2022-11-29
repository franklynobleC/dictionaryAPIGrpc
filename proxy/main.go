package main

import (
	"context"
	"log"

	service "github.com/franklynobleC/dictionaryAPIGrpc/proto/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	// "google.golang.org/api/transport/grpc"
	// "google.golang.org/grpc"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
)

func main() {

	mux := runtime.NewServeMux()

	err := service.RegisterEnglishDictionaryHandlerServer (context.Background(), mux, "localhost:5000", []grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		log.Fatalf("dictionary registration server Error %v", err)
	}

}
