package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// Delete ...
func (s *StorageServer) Delete(ctx context.Context, req *pb.DeleteReq) (*pb.DeleteRes, error) {
	var err = s.s.Delete(req.GetID())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteRes{}, nil
}
