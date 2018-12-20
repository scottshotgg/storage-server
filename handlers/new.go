package handlers

// StorageServer ...
type StorageServer struct{}

// New ...
func New() (*StorageServer, error) {
	return &StorageServer{}, nil
}
