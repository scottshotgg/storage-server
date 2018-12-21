package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// Get ...
func (s *StorageServer) Get(ctx context.Context, req *pb.GetReq) (*pb.GetRes, error) {
	var item, err = s.s.Get(ctx, req.GetItemID())
	if err != nil {
		return nil, err
	}

	if item == nil {
		// return 404
	}

	return &pb.GetRes{
		Item: item.ToProto(),
	}, nil
}
