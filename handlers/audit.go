package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// Audit ...
func (s *StorageServer) Audit(ctx context.Context, in *pb.AuditReq) (*pb.AuditRes, error) {
	return nil, nil
}
