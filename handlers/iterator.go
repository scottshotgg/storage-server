package handlers

import (
	pb "github.com/scottshotgg/storage/protobufs"
	"github.com/scottshotgg/storage/storage"
	"google.golang.org/api/iterator"
)

// Iterator ...
func (s *StorageServer) Iterator(req *pb.IteratorReq, server pb.Storage_IteratorServer) error {
	var iter, err = s.s.Iterator()
	if err != nil {
		return err
	}

	var (
		items   = make([]storage.Item, chunkAmount)
		sendErr error
	)

	for {
		// Chunk the messages to `chunkAmount`
		for i := 0; i < chunkAmount; i++ {
			items[i], err = iter.Next()
			if err != nil {
				if err == iterator.Done {
					break
				}

				// Do we just return here or skip?
				return err
			}
		}

		sendErr = server.Send(&pb.IteratorRes{
			Items: pbItemsFromItems(items),
		})
		if sendErr != nil {
			return sendErr
		}

		if err == iterator.Done {
			break
		}
	}

	return nil
}
