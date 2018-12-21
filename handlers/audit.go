package handlers

import (
	"context"

	pb "github.com/scottshotgg/storage/protobufs"
)

// Audit ...
func (s *StorageServer) Audit(ctx context.Context, req *pb.AuditReq) (*pb.AuditRes, error) {
	var clMap, err = s.s.Audit()
	if err != nil {
		return nil, err
	}

	return &pb.AuditRes{
		Changelogs: pbChangelogMapFromMap(clMap),
	}, nil
}
