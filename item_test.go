package gomediacenter

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var matchID = NewID()
var missingItem = NewID()

func TestStaticProvider(t *testing.T) {
	assert := assert.New(t)
	// Setup
	p1 := createMockedStaticProvider("p1", nil, true)
	p2 := createMockedStaticProvider("p2", nil, false)
	p3 := createMockedStaticProvider("p3", ErrItemMissing, true)
	p4 := createMockedStaticProvider("p4", nil, false)
	RegisterStaticProvider(p1)
	RegisterStaticProvider(p2)
	RegisterStaticProvider(p3)
	RegisterStaticProvider(p4)

	// Tests
	t.Run("StaticProviderDomain", func(t *testing.T) {
		t.Run("ShouldNotAddRegisteredProvider", func(t *testing.T) {
			t.Parallel()
			err := RegisterStaticProvider(p1)
			assert.Equal(ErrProviderAlreadyRegistered, err)
		})
		t.Run("CanRemoveProvider", func(t *testing.T) {
			t.Parallel()
			err := RemoveStaticProvider(p4)
			assert.NoError(err)
		})
		t.Run("ShouldNotRemoveNonRegisteredProvider", func(t *testing.T) {
			t.Parallel()
			p := createMockedStaticProvider("p", nil, false)
			err := RemoveStaticProvider(p)
			assert.Error(err)
			assert.Equal(ErrProviderNotRegistered, err)
		})
		t.Run("ShouldReturnErrorWhenItemIsMissing", func(t *testing.T) {
			t.Parallel()
			id := NewID()
			r, err := GetItem(id)
			assert.Equal(ErrNoProviderFound, err)
			assert.Nil(r)
		})
		t.Run("ShouldReturnItemMissing", func(t *testing.T) {
			t.Parallel()
			r, err := GetItem(missingItem)
			assert.Equal(ErrItemMissing, err)
			assert.Nil(r)
		})
		t.Run("ShouldReturnItem", func(t *testing.T) {
			t.Parallel()
			r, err := GetItem(matchID)
			assert.NoError(err)
			b, err := ioutil.ReadAll(r)
			assert.NoError(err)
			assert.Equal("p1", string(b))
		})
	})
	// Teardown
	RemoveStaticProvider(p1)
	RemoveStaticProvider(p2)
	RemoveStaticProvider(p3)
	RemoveStaticProvider(p4)

	// Setup
	var ErrFromStaticProvider = errors.New("Error from provider")
	p := createMockedStaticProvider("p", ErrFromStaticProvider, true)

	// Test
	t.Run("StaticProviderInstance", func(t *testing.T) {
		RegisterStaticProvider(p)
		r, err := GetItem(matchID)
		assert.Nil(r)
		assert.Error(err)
		assert.Equal(ErrFromStaticProvider, err)
	})

	// Teardown
	RemoveStaticProvider(p)
}

func createMockedStaticProvider(item string, err error, provides bool) *staticProvider {
	id := NewID()
	p := &staticProvider{
		id: id,
		doProvides: func(id ID) bool {
			if id.Equal(matchID) {
				return provides
			} else if id.Equal(missingItem) && item == "p3" {
				return provides
			}
			return false
		},
		doGetItem: func(id ID) (ReadSeekCloser, error) {
			if id.Equal(matchID) {
				b := bytes.NewReader([]byte(item))
				return &itemMock{Name: item, buf: b}, err
			} else if id.Equal(missingItem) {
				return nil, err
			}
			return nil, err
		},
		doGetID: func() ID { return id },
	}
	return p
}

type itemMock struct {
	Name string
	buf  io.Reader
}

func (i *itemMock) Read(p []byte) (n int, err error) {
	return i.buf.Read(p)
}

func (i *itemMock) Seek(offset int64, whence int) (int64, error) {
	panic("not implemented")
}

func (i *itemMock) Close() error {
	panic("not implemented")
}

type staticProvider struct {
	doGetItem  func(id ID) (ReadSeekCloser, error)
	doProvides func(id ID) bool
	doGetID    func() ID
	id         ID
}

func (p *staticProvider) ProvideItem(id ID) bool {
	return p.doProvides(id)
}

func (p *staticProvider) GetItem(id ID) (ReadSeekCloser, error) {
	return p.doGetItem(id)
}

func (p *staticProvider) GetID() ID {
	return p.id
}
