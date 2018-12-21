package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// GetBy ...
func (s *StorageServer) GetBy(ctx context.Context, req *pb.GetByReq) (*pb.GetByRes, error) {
	// TODO: change this to int64
	var items, err = s.s.GetBy(ctx, req.GetKey(), req.GetOp(), req.GetValue(), int(req.GetLimit()))
	if err != nil {
		return nil, err
	}

	if items == nil {
		// return 404
	}

	return &pb.GetByRes{
		Items: pbItemsFromItems(items),
	}, nil
}
