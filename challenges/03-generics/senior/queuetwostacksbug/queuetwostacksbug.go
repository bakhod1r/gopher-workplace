// Package queuetwostacksbug — Gopher Workplace challenge.
package queuetwostacksbug

// Queue is a FIFO queue backed by two stacks.
type Queue[T any] struct {
	in  []T
	out []T
}

// Dequeue removes and returns the oldest element and true.
// It returns the zero value and false when empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	for len(q.in) > 0 {
		v := q.in[len(q.in)-1]
		q.in = q.in[:len(q.in)-1]
		q.out = append(q.out, v)
	}
	if len(q.out) == 0 {
		var zero T
		return zero, false
	}
	v := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return v, true
	// CHANGE CODE ABOVE THIS LINE
}

// Enqueue adds v to the back of the queue.
func (q *Queue[T]) Enqueue(v T) {
	q.in = append(q.in, v)
}
