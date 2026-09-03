// Package ringbuf — Gopher Workplace challenge.
package ringbuf

// Ring is a fixed-capacity circular buffer of ints.
type Ring struct {
	buf  []int
	head int
	n    int
}

// NewRing returns a ring that holds at most cap elements.
func NewRing(cap int) *Ring {
	if cap < 1 {
		cap = 1
	}
	return &Ring{buf: make([]int, cap)}
}

// Len reports how many elements the ring currently holds.
func (r *Ring) Len() int { return r.n }

// Items returns the ring's contents from oldest to newest.
func (r *Ring) Items() []int {
	out := make([]int, 0, r.n)
	for i := 0; i < r.n; i++ {
		out = append(out, r.buf[(r.head+i)%len(r.buf)])
	}
	return out
}

// Push adds v to the ring, overwriting the oldest element once the ring
// is full.
//
// The ring never grows: it was given its capacity at construction and every
// later Push must reuse that storage.
//
// Examples:
//
//	r := NewRing(2); r.Push(1); r.Push(2); r.Push(3) => Items() is [2 3]
func (r *Ring) Push(v int) {
	panic("not implemented")
}
