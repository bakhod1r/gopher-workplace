// Package pqgen — Gopher Workplace challenge.
package pqgen

// entry pairs a value with its scheduling priority.
type entry[T any] struct {
	value    T
	priority int
	seq      int
}

// PQ is a priority queue of T. Lower priority numbers run first.
// Its zero value is an empty queue.
type PQ[T any] struct {
	items []entry[T]
	seq   int
}

// Push adds v with the given priority. Lower runs first.
func (q *PQ[T]) Push(v T, priority int) {
	// TODO(candidate): store the item, keeping the queue ordered by priority.
	panic("not implemented")
}

// Pop removes and returns the highest-priority item and true.
// Items of equal priority come out in insertion order.
func (q *PQ[T]) Pop() (T, bool) {
	// TODO(candidate): remove and return the front item.
	panic("not implemented")
}

// Len returns the number of queued items.
func (q *PQ[T]) Len() int {
	// TODO(candidate): report how many items are queued.
	panic("not implemented")
}
