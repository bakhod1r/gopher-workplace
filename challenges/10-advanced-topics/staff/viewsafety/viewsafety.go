// Package viewsafety — Gopher Workplace challenge.
package viewsafety

import "unsafe"

// Window returns the n bytes of b starting at off, as a view whose
// capacity is exactly n.
//
// The caller may append to the result, so the capacity must not let that
// append reach the bytes after the window.
//
// Examples:
//
//	Window(buf, 2, 3) => buf[2:5] with capacity 3, true
func Window(b []byte, off, n int) ([]byte, bool) {
	panic("not implemented")
}
