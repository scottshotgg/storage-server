package handlers

import (
	"context"

	"github.com/scottshotgg/storage/object"
	pb "github.com/scottshotgg/storage/protobufs"
)

// Set ...
func (s *StorageServer) Set(ctx context.Context, req *pb.SetReq) (*pb.SetRes, error) {
	var item = req.GetItem()
	if item == nil {
		// return 404
	}

	// TODO: change this later
	var err = s.s.Set(ctx, object.FromProto(item))
	if err != nil {
		return nil, err
	}

	return &pb.SetRes{}, nil
}
