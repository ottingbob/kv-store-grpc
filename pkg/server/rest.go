package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/ottingbob/kv-store-grpc/pkg/gen"
	"google.golang.org/grpc"
)

func RunREST() error {
	fmt.Println("Key-Value REST Service Initializing...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		gwmux              = runtime.NewServeMux()
		grpcServerEndpoint = "localhost:5000"
		opts               = []grpc.DialOption{grpc.WithInsecure()}
	)

	err := pb.RegisterKVServiceHandlerFromEndpoint(ctx, gwmux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/", gwmux)
	// serve swagger endpoint
	return http.ListenAndServe(":8080", mux)
}
