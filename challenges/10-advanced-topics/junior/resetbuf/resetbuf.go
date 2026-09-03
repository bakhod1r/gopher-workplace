// Package resetbuf — Gopher Workplace challenge.
package resetbuf

// Reset returns buf emptied for reuse, keeping the capacity it already
// has so the next writer does not have to allocate again.
//
// Examples:
//
//	Reset(make([]byte, 8, 64)) => length 0, capacity 64
func Reset(buf []byte) []byte {
	panic("not implemented")
}
