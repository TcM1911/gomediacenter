package gomediacenter

import (
	"encoding/hex"
	"errors"

	uuid "github.com/satori/go.uuid"
)

// ID is a unique identifier for an object. A new ID
// can be generated by using the New function. For more
// secure random ID, use the NewRandom function.
type ID struct {
	uuid uuid.UUID
}

// NewID generates a new ID.
func NewID() ID {
	return ID{uuid: uuid.NewV1()}
}

// NewRandomID generates a more secure random ID that can be used
// for session tokens.
func NewRandomID() ID {
	return ID{uuid: uuid.NewV4()}
}

// IDFromBytes converts a byte slice to an ID.
func IDFromBytes(b []byte) (ID, error) {
	id := ID{}
	if len(b) != 16 {
		return id, errors.New("byte array has to have a length of 16")
	}
	uid, err := uuid.FromBytes(b)
	if err != nil {
		return id, err
	}
	id.uuid = uid
	return id, nil
}

// IDFromString converts a string into an ID.
func IDFromString(s string) (ID, error) {
	if len(s) != 32 {
		return ID{}, errors.New("the id string must be 32 characters long")
	}
	buf := make([]byte, 16)
	hex.Decode(buf, []byte(s))
	return IDFromBytes(buf)
}

// String returns a string representation of the ID.
func (i *ID) String() string {
	buf := make([]byte, 32)
	hex.Encode(buf, i.uuid.Bytes())
	return string(buf)
}

// Equal compares the ID to ID b and returns true if
// they are equal.
func (i *ID) Equal(b *ID) bool {
	return uuid.Equal(i.uuid, b.uuid)
}
