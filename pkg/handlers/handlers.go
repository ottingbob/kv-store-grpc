package handlers

import (
	"context"
	"fmt"

	pb "github.com/ottingbob/kv-store-grpc/pkg/gen"
	// "go.uber.org/zap"
	// "go.uber.org/zap"
	// "gorm.io/gorm"
	// "github.com/ottingbob/quick-bufs/pkg/database"
	// "go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	// "gorm.io/gorm"
)

// Simple in memory database to store string key-value pairs.
// Normally we would make a more complex model with a unique identifier in addition to a user provided slug of some type
// but this is the simple solution
var database = map[string]string{}

type KVServiceRPCServer struct {
	Name string
	// QSController *quickServiceController
	pb.UnimplementedKVServiceServer
	// DB           *gorm.DB
	// Log          *zap.Logger
}

func (s KVServiceRPCServer) CreateKV(ctx context.Context, msg *pb.KVMessage) (*pb.KVMessage, error) {
	fmt.Println("Hitting CreateKV RPC Endpoint...", msg)
	database[msg.Key] = msg.Value
	return msg, nil
}

func (s KVServiceRPCServer) DeleteKV(ctx context.Context, req *pb.DeleteKVRequest) (*emptypb.Empty, error) {
	fmt.Println("Hitting DeleteKV RPC Endpoint...")
	_, ok := database[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key-Value pair could not be found at key %s", req.Key)
	}
	delete(database, req.Key)
	return &emptypb.Empty{}, nil
}

func (s KVServiceRPCServer) GetKV(ctx context.Context, req *pb.GetKVRequest) (*pb.KVMessage, error) {
	fmt.Println("Hitting GetKV RPC Endpoint...")
	dbValue, ok := database[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key-Value pair could not be found at key %s", req.Key)
	}
	return &pb.KVMessage{
		Key:   req.Key,
		Value: dbValue,
	}, nil
}
