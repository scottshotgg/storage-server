package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// Sync ...
func (s *StorageServer) Sync(ctx context.Context, req *pb.SyncReq) (*pb.SyncRes, error) {
	var err = s.s.Sync()
	if err != nil {
		return nil, err
	}

	return &pb.SyncRes{}, nil
}
