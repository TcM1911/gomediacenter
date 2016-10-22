package gomediacenter

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"

	uuid "github.com/satori/go.uuid"
)

// ID is a unique identifier for an object. A new ID
// can be generated by using the New function. For more
// secure random ID, use the NewRandom function.
type ID struct {
	Bytes []byte
}

// NewID generates a new ID.
func NewID() ID {
	return ID{Bytes: uuid.NewV1().Bytes()}
}

// NewRandomID generates a more secure random ID that can be used
// for session tokens.
func NewRandomID() ID {
	return ID{Bytes: uuid.NewV4().Bytes()}
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
	id.Bytes = uid.Bytes()
	return id, nil
}

// IDFromString converts a string into an ID.
func IDFromString(s string) (ID, error) {
	if len(s) != 32 {
		return ID{}, errors.New("the id string must be 32 characters long")
	}
	buf := make([]byte, 16)
	_, err := hex.Decode(buf, []byte(s))
	if err != nil {
		return ID{}, err
	}
	return IDFromBytes(buf)
}

// String returns a string representation of the ID.
func (i *ID) String() string {
	buf := make([]byte, 32)
	hex.Encode(buf, i.Bytes)
	return string(buf)
}

// Equal compares the ID to ID b and returns true if
// they are equal.
func (i *ID) Equal(b ID) bool {
	return bytes.Equal(i.Bytes, b.Bytes)
}

// IsNil checks if the ID is nil.
func (i *ID) IsNil() bool {
	nullUUID := uuid.NullUUID{}
	nullID := ID{Bytes: nullUUID.UUID.Bytes()}
	return i.Equal(nullID)
}

// GetIDFromContext gets the ID from the context. If no ID exist in the context,
// a null byte ID is returned.
func GetIDFromContext(ctx context.Context) ID {
	id, ok := ctx.Value(ctxKey).(ID)
	if !ok {
		null := uuid.NullUUID{}
		return ID{Bytes: null.UUID.Bytes()}
	}
	return id
}

// AddIDToContext adds the ID to the context.
func AddIDToContext(ctx context.Context, id ID) context.Context {
	return context.WithValue(ctx, ctxKey, id)
}

const ctxKey string = "idKey"
