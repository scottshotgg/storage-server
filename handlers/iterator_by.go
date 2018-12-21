package handlers

import (
	pb "github.com/scottshotgg/storage/protobufs"
	"github.com/scottshotgg/storage/storage"
	"google.golang.org/api/iterator"
)

// IteratorBy ...
func (s *StorageServer) IteratorBy(req *pb.IteratorByReq, server pb.Storage_IteratorByServer) error {
	var iter, err = s.s.IteratorBy(req.GetKey(), req.GetOp(), req.GetValue())
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

		sendErr = server.Send(&pb.IteratorByRes{
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
