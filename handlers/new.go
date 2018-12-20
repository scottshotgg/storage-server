package handlers

import (
	dberrors "github.com/scottshotgg/storage/errors"
	"github.com/scottshotgg/storage/storage"
)

// StorageServer ...
type StorageServer struct {
	s storage.Storage
}

// New creates a new StorageServer with from a storage.Storage
func New(node storage.Storage) (*StorageServer, error) {
	if node == nil {
		return nil, dberrors.ErrNilDB
	}

	return &StorageServer{
		s: node,
	}, nil
}
