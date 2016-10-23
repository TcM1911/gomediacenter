package gomediacenter

import (
	"context"
	"encoding/hex"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := NewID()
	expected := ID{}
	assert.IsType(t, expected, id)
}

func TestNewRandomID(t *testing.T) {
	id := NewRandomID()
	expected := ID{}
	assert.IsType(t, expected, id)
}

func TestIDStringer(t *testing.T) {
	id := NewID()
	assert.IsType(t, "", id.String())
}

func TestIDFromString(t *testing.T) {
	assert := assert.New(t)
	id, err := IDFromString(idStr)
	assert.NoError(err, "Should not return an error.")
	assert.IsType(ID{}, id, "Should be an ID type.")
	assert.Equal(idStr, id.String(), "Should return expected ID string.")
}

func TestIDFromBytes(t *testing.T) {
	assert := assert.New(t)
	b := make([]byte, 16)
	hex.Decode(b, []byte(idStr))
	id, err := IDFromBytes(b)
	assert.NoError(err, "Should not return an error.")
	assert.IsType(ID{}, id, "Should be an ID type.")
	assert.Equal(idStr, id.String(), "Should return expected ID string.")
}

func TestIDFromBytesToLong(t *testing.T) {
	assert := assert.New(t)
	toLong := "00112233445566778899aabbccddeeffaa"
	b := make([]byte, 17)
	hex.Decode(b, []byte(toLong))
	_, err := IDFromBytes(b)
	assert.Error(err, "Should return an error.")
}

func TestIDFromStringToLong(t *testing.T) {
	assert := assert.New(t)
	toLong := "00112233445566778899aabbccddeeffaa"
	_, err := IDFromString(toLong)
	assert.Error(err, "Should return an error.")
}

func TestIDEqual(t *testing.T) {
	assert := assert.New(t)
	a := NewID()
	b := NewRandomID()
	c := NewRandomID()
	e, err := IDFromString(idStr)
	assert.NoError(err)
	f, err := IDFromString(idStr)
	assert.NoError(err)
	assert.False(a.Equal(b))
	assert.False(a.Equal(c))
	assert.False(b.Equal(c))
	assert.True(e.Equal(f))
}

func TestIsNil(t *testing.T) {
	assert := assert.New(t)
	id := NewID()
	assert.False(id.IsNil(), "Should return false for non null id.")

	nullUUID := uuid.NullUUID{}
	nullID := ID{Bytes: nullUUID.UUID.Bytes()}
	assert.True(nullID.IsNil(), "Should return true for null id.")
}

func TestContext(t *testing.T) {
	assert := assert.New(t)
	t.Run("Retrieving from context", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		nul := GetIDFromContext(ctx)
		assert.True(nul.IsNil(), "Should return a null id")
		id := NewID()
		ctx = context.WithValue(ctx, idCtxKey, id)
		assert.Equal(GetIDFromContext(ctx), id, "Should return the id")
	})
	t.Run("Saving to context", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		id := NewID()
		ctx = AddIDToContext(ctx, id)
		actual := ctx.Value(idCtxKey).(ID)
		assert.Equal(id, actual, "Should add id to context")
	})
}

const idStr = "00112233445566778899aabbccddeeff"
