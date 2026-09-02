// Package dequewrapmodbug — Gopher Workplace challenge.
package dequewrapmodbug

// Deque is a double-ended queue over a ring buffer.
type Deque[T any] struct {
	buf  []T
	head int
	n    int
}

// PushFront inserts v at the front of the deque.
func (d *Deque[T]) PushFront(v T) {
	// CHANGE CODE BELOW THIS LINE
	d.grow()
	d.head = (d.head - 1) % len(d.buf)
	d.buf[d.head] = v
	d.n++
	// CHANGE CODE ABOVE THIS LINE
}

// grow makes room for one more element, re-laying the ring out from 0.
func (d *Deque[T]) grow() {
	if d.n < len(d.buf) {
		return
	}
	nb := make([]T, 2*len(d.buf)+1)
	for i := 0; i < d.n; i++ {
		nb[i] = d.buf[(d.head+i)%len(d.buf)]
	}
	d.buf = nb
	d.head = 0
}

// PushBack appends v at the back of the deque.
func (d *Deque[T]) PushBack(v T) {
	d.grow()
	d.buf[(d.head+d.n)%len(d.buf)] = v
	d.n++
}

// PopFront removes and returns the front element and true.
// It returns the zero value and false when the deque is empty.
func (d *Deque[T]) PopFront() (T, bool) {
	if d.n == 0 {
		var zero T
		return zero, false
	}
	v := d.buf[d.head]
	d.head = (d.head + 1) % len(d.buf)
	d.n--
	return v, true
}

// At returns the i-th element from the front.
func (d *Deque[T]) At(i int) (T, bool) {
	if i < 0 || i >= d.n {
		var zero T
		return zero, false
	}
	return d.buf[(d.head+i)%len(d.buf)], true
}

// Len reports how many elements the deque holds.
func (d *Deque[T]) Len() int {
	return d.n
}
