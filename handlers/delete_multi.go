package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// TODO: Should really be sending back a status for each one
// TODO: make this

// DeleteMulti ...
func (s *StorageServer) DeleteMulti(ctx context.Context, req *pb.DeleteReq) (*pb.DeleteRes, error) {
	// var items, err = s.s.DeleteMulti(ctx, req.GetIDs())
	// if err != nil {
	// 	return nil, err
	// }

	// if items == nil {
	// 	// return 404
	// }

	// return &pb.DeleteMultiRes{
	// 	Items: pbItemsFromItems(items),
	// }, nil

	return nil, nil
}
