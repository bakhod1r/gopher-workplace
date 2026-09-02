// Package nodegen — Gopher Workplace challenge.
package nodegen

// Node is one element of a singly linked list of T.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// Prepend returns a new list with v at the front.
func Prepend[T any](head *Node[T], v T) *Node[T] {
	// TODO(candidate): return a node pointing at the old head.
	panic("not implemented")
}

// ToSlice returns the values from head to tail.
func ToSlice[T any](head *Node[T]) []T {
	// TODO(candidate): walk the chain, collecting values.
	panic("not implemented")
}
