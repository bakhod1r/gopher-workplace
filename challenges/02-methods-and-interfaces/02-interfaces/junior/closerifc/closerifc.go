// Package closerifc — Gopher Workplace challenge.
package closerifc

import "errors"

// ErrAlreadyClosed is returned by a second Close.
var ErrAlreadyClosed = errors.New("already closed")

// Closer releases a resource.
type Closer interface {
	Close() error
}

// File is a fake file handle.
type File struct {
	Closed bool
}

// Close releases the file.
//
// Examples:
//
//	f := &File{}; f.Close() => nil
//	f.Close()               => ErrAlreadyClosed
func (f *File) Close() error {
	// TODO(candidate): reject a double close, otherwise mark closed.
	panic("not implemented")
}

// CloseAll closes every closer and returns the first error seen.
func CloseAll(cs []Closer) error {
	// TODO(candidate): close each one, stop at the first error.
	panic("not implemented")
}
