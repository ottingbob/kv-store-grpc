package handlers

import (
	"context"

	pb "github.com/ottingbob/kv-store-grpc/pkg/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// Simple in memory database to store string key-value pairs.
// Normally we would make a more complex model with a unique identifier in addition to a user provided slug of some type
// but this is the simple solution
var database = map[string]string{}

type KVServiceRPCServer struct {
	pb.UnimplementedKVServiceServer
}

func (s KVServiceRPCServer) CreateKV(ctx context.Context, msg *pb.KVMessage) (*pb.KVMessage, error) {
	database[msg.Key] = msg.Value
	return msg, nil
}

func (s KVServiceRPCServer) DeleteKV(ctx context.Context, req *pb.DeleteKVRequest) (*emptypb.Empty, error) {
	_, ok := database[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key-Value pair could not be found at key %s", req.Key)
	}
	delete(database, req.Key)
	return &emptypb.Empty{}, nil
}

func (s KVServiceRPCServer) GetKV(ctx context.Context, req *pb.GetKVRequest) (*pb.KVMessage, error) {
	dbValue, ok := database[req.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Key-Value pair could not be found at key %s", req.Key)
	}
	return &pb.KVMessage{
		Key:   req.Key,
		Value: dbValue,
	}, nil
}
