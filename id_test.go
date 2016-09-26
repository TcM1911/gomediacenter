package gomediacenter

import (
	"encoding/hex"
	"testing"

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

const idStr = "00112233445566778899aabbccddeeff"
