// Package zerocopykeys — Gopher Workplace challenge.
package zerocopykeys

import "unsafe"

// Count increments m's counter for each key.
//
// The lookup may borrow the key's bytes, but a key that ends up stored in
// the map must own its bytes: the caller reuses the buffers the keys point
// into.
//
// Examples:
//
//	Count(m, [][]byte{[]byte("a")}) => m["a"] == 1
func Count(m map[string]int, keys [][]byte) {
	panic("not implemented")
}
