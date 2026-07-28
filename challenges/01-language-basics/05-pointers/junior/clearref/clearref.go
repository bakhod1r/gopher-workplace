// Package clearref — Gopher Workplace challenge.
package clearref

type Node struct {
	Val  int
	Next *Node
}

// Detach sets the node's Next to nil, unlinking the rest of the list.
func Detach(n *Node) {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
