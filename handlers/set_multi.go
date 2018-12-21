package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// TODO: should really be returning a status for each one

// SetMulti ...
func (s *StorageServer) SetMulti(ctx context.Context, req *pb.SetMultiReq) (*pb.SetMultiRes, error) {
	var err = s.s.SetMulti(ctx, protoToItems(req.GetItems()))
	if err != nil {
		return nil, err
	}

	return &pb.SetMultiRes{}, nil
}
