// Package arena — Gopher Workplace challenge.
package arena

// Arena hands out blocks of one fixed backing allocation.
type Arena struct {
	buf  []byte
	used uintptr
}

// NewArena returns an arena of size bytes.
func NewArena(size int) *Arena {
	if size < 0 {
		size = 0
	}
	return &Arena{buf: make([]byte, size)}
}

// Used reports how many bytes have been handed out, including padding.
func (a *Arena) Used() int { return int(a.used) }

// Cap reports the arena's total size.
func (a *Arena) Cap() int { return len(a.buf) }

// Alloc returns the next n bytes of the arena, starting at an offset that
// is a multiple of align.
//
// The arena never grows: when the remaining space cannot satisfy the
// request, Alloc reports false.
//
// Examples:
//
//	a := NewArena(64); a.Alloc(8, 8) => an 8-byte block, true
func (a *Arena) Alloc(n int, align uintptr) ([]byte, bool) {
	panic("not implemented")
}
