package gomediacenter

import (
	"io"
)

// ReadSeekCloser describes the interface for ReadSeekCloser. It adds the io.Closer
// interface to the io.ReadSeeker interface.
type ReadSeekCloser interface {
	io.ReadSeeker
	io.Closer
}
