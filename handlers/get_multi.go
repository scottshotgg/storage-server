package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// GetMulti ...
func (s *StorageServer) GetMulti(ctx context.Context, req *pb.GetMultiReq) (*pb.GetMultiRes, error) {
	var items, err = s.s.GetMulti(ctx, req.GetIDs())
	if err != nil {
		return nil, err
	}

	if items == nil {
		// return 404
	}

	return &pb.GetMultiRes{
		Items: pbItemsFromItems(items),
	}, nil
}
