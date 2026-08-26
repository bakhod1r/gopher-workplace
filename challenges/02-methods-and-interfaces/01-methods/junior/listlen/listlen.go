// Package listlen — Gopher Workplace challenge.
package listlen

// Node is a singly-linked list node.
type Node struct {
	Val  int
	Next *Node
}

// Len returns the number of nodes in the list starting from this node.
//
// Examples:
//
//	node(1 -> 2 -> 3).Len() => 3
//	(*Node)(nil).Len()       => 0
func (n *Node) Len() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
