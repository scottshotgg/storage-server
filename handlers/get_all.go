package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

const (
	chunkAmount = 100
)

// TODO: i dont think this is a smart functionality to even provide; would be easy to abuse even if inadvertently

// GetAll ...
func (s *StorageServer) GetAll(ctx context.Context, req *pb.GetAllReq) (*pb.GetAllRes, error) {
	return &pb.GetAllRes{}, nil
}
