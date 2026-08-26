// Package iterpatt — Gopher Workplace challenge.
package iterpatt

// Node is a linked list node.
type Node struct {
	Val  int
	Next *Node
}

// ListIter iterates over a linked list.
type ListIter struct {
	current *Node
}

// NewIter creates an iterator starting at head.
func NewIter(head *Node) *ListIter {
	return &ListIter{current: head}
}

// HasNext returns true if there is a current node.
func (it *ListIter) HasNext() bool {
	// TODO(candidate): return current != nil
	panic("not implemented")
}

// Next returns the current value and advances the iterator.
// Assumes HasNext() is true.
func (it *ListIter) Next() int {
	// TODO(candidate): save current.Val, advance current to current.Next, return saved val
	panic("not implemented")
}
