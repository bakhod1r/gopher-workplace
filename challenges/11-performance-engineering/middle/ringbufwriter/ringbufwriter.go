// Package ringbufwriter — Gopher Workplace challenge.
package ringbufwriter

// Ring keeps the most recent n values in a fixed array, overwriting the
// oldest — the shape every "last 1000 log lines" buffer has. It never
// allocates after construction.
type Ring struct {
	buf   []int
	next  int
	size  int
	count int64
}

// New returns a ring holding at most n values. A non-positive n holds nothing.
//
// Examples:
//
//	New(3)
func New(n int) *Ring {
	panic("not implemented")
}

// Add stores v, overwriting the oldest value once the ring is full.
//
// Examples:
//
//	r.Add(1)
func (r *Ring) Add(v int) {
	panic("not implemented")
}

// Len returns how many values the ring currently holds, and Total how many
// have ever been added.
//
// Examples:
//
//	r.Len() => 3
func (r *Ring) Len() int { panic("not implemented") }

// Total reports how many values have been added over the ring's lifetime,
// including those since overwritten.
//
// Examples:
//
//	r.Total() => 10
func (r *Ring) Total() int64 { panic("not implemented") }

// Snapshot returns the held values oldest first, in a fresh slice.
//
// Examples:
//
//	New(3) with 1,2,3,4 added => []int{2, 3, 4}
func (r *Ring) Snapshot() []int {
	panic("not implemented")
}
