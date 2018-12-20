package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// Sync ...
func (s *StorageServer) Sync(ctx context.Context, in *pb.SyncReq) (*pb.SyncRes, error) {
	return nil, nil
}
