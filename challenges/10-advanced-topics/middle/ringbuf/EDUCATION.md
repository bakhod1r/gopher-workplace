# A Buffer That Allocates Once, Ever

## Intuition

A ring turns "keep the last N" into pure index arithmetic. Nothing is ever appended, moved or freed — the write position simply walks around a fixed array.

## Approach

1. If `r.n < len(r.buf)`, write at `(head+n) % len(buf)` and increment `n`.
2. Otherwise overwrite `buf[head]` and advance `head` by one, modulo the capacity.

## Solution

```go
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
// 	r := NewRing(2); r.Push(1); r.Push(2); r.Push(3) => Items() is [2 3]
func (r *Ring) Push(v int) {
	if r.n < len(r.buf) {
		r.buf[(r.head+r.n)%len(r.buf)] = v
		r.n++
		return
	}
	r.buf[r.head] = v
	r.head = (r.head + 1) % len(r.buf)
}
```

## Walkthrough

Capacity 3, pushes 1,2,3 fill indices 0,1,2 with head 0. Push 4 overwrites index 0 and moves head to 1, so the oldest is now index 1 — [2 3 4].

## Pitfalls

- Incrementing `n` past the capacity — the length is capped once the ring is full.
- Using `append` for the not-yet-full case; the buffer already has the slots.
