package handlers

import (
	"github.com/scottshotgg/storage/object"
	pb "github.com/scottshotgg/storage/protobufs"
	"github.com/scottshotgg/storage/storage"
)

func pbItemsFromItems(items []storage.Item) []*pb.Item {
	var pbItems = make([]*pb.Item, len(items))

	for i := range items {
		pbItems[i] = items[i].ToProto()
	}

	return pbItems
}

func pbChangelogFromChangelog(cl *storage.Changelog) *pb.Changelog {
	return &pb.Changelog{
		ID:        cl.ID,
		Timestamp: cl.Timestamp,
		ItemID:    cl.ObjectID,
	}
}

func pbChangelogMapFromMap(cls map[string]*storage.Changelog) map[string]*pb.Changelog {
	var pbCLMap = map[string]*pb.Changelog{}

	for id, cl := range cls {
		pbCLMap[id] = pbChangelogFromChangelog(cl)
	}

	return pbCLMap
}

func protoToItems(pbItems []*pb.Item) []storage.Item {
	var items = make([]storage.Item, len(pbItems))

	for i := range pbItems {
		// TODO: fix this later
		items[i] = object.FromProto(pbItems[i])
	}

	return items
}
