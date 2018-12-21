package storage

import (
	pb "github.com/scottshotgg/storage/protobufs"
)

// TODO: might try doing this later and using this for the Object Value
type Value interface {
	// Require proto marshalers
	Marshal() (data []byte, err error)
	Unmarshal(data []byte) error
}

// type Key interface {
// 	Type() keyType
// 	Value() interface{}
// }

// These are named `Get<prop>()` so that the Item interface
// and protobuf Item can satisfy each other

type Item interface {
	// Properties
	GetID() string
	GetValue() []byte
	GetTimestamp() int64
	GetKeys() []string
	GetDeleted() bool

	// These are needed mainly for Redis
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error

	// Require proto marshalers
	Marshal() (data []byte, err error)
	Unmarshal(data []byte) error

	// // Require Gob encoding/decoding to be implemented
	// GobEncode() ([]byte, error)
	// GobDecode([]byte) error

	ToProto() *pb.Item
}
