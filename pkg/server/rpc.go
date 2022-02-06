package server

import (
	"fmt"
	"net"

	gw "github.com/ottingbob/kv-store-grpc/pkg/gen"
	"github.com/ottingbob/kv-store-grpc/pkg/handlers"

	// "github.com/ottingbob/kv-store-grpc/pkg/database"
	// "go.uber.org/zap"

	"google.golang.org/grpc"
)

func RunKVServiceRPC() error {
	/*
		log, err := zap.NewProduction()
		if err != nil {
			return err
		}
		defer log.Sync()
		log.Info("quick service rpc cmd started")
	*/
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

	/*
		db, err := database.Open()
		if err != nil {
			return err
		}
	*/

	rpcServer := grpc.NewServer()
	kvRPCServer := &handlers.KVServiceRPCServer{
		Name: "Ready-Freddy",
		// DB:           db,
		// Log:          log,
	}
	gw.RegisterKVServiceServer(rpcServer, kvRPCServer)
	return rpcServer.Serve(lis)
}
