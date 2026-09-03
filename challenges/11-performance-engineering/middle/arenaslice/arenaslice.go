// Package arenaslice — Gopher Workplace challenge.
package arenaslice

// Arena hands out slices carved from one backing block: a bump allocator.
// Everything it returns dies together when Reset is called, which is what
// makes it fast — no per-object bookkeeping and no per-object free.
type Arena struct {
	block []byte
	used  int
}

// NewArena returns an arena backed by a block of n bytes. A non-positive n
// gives an arena that satisfies nothing.
//
// Examples:
//
//	NewArena(1024)
func NewArena(n int) *Arena {
	panic("not implemented")
}

// Alloc carves n zeroed bytes out of the block and reports whether it fit.
// A failed allocation must not consume any of the block. A non-positive n
// returns an empty, non-nil slice and true.
//
// Examples:
//
//	a.Alloc(16) => 16 bytes, true
func (a *Arena) Alloc(n int) ([]byte, bool) {
	panic("not implemented")
}

// Used and Free report the block's occupancy.
//
// Examples:
//
//	a.Used() => 16
func (a *Arena) Used() int { panic("not implemented") }

// Free reports how many bytes of the block remain.
//
// Examples:
//
//	NewArena(1024).Free() => 1024
func (a *Arena) Free() int { panic("not implemented") }

// Reset frees everything the arena handed out at once. Previously returned
// slices still point into the block and must not be used afterwards.
//
// Examples:
//
//	a.Reset(); a.Used() == 0
func (a *Arena) Reset() {
	panic("not implemented")
}
