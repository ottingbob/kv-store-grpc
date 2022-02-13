package server

import (
	"fmt"
	"net"

	gw "github.com/ottingbob/kv-store-grpc/pkg/gen"
	"github.com/ottingbob/kv-store-grpc/pkg/handlers"

	"google.golang.org/grpc"
)

func RunKVServiceRPC() error {
	fmt.Println("Key-Value RPC Service Initializing...")

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}

	defer func() {
		err = lis.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	rpcServer := grpc.NewServer()
	kvRPCServer := &handlers.KVServiceRPCServer{}
	gw.RegisterKVServiceServer(rpcServer, kvRPCServer)
	return rpcServer.Serve(lis)
}
