package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// QuickSync ...
func (s *StorageServer) QuickSync(ctx context.Context, req *pb.QuickSyncReq) (*pb.QuickSyncRes, error) {
	var err = s.s.QuickSync()
	if err != nil {
		return nil, err
	}

	return &pb.QuickSyncRes{}, nil
}
