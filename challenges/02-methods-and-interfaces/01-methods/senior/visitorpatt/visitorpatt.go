// Package visitorpatt — Gopher Workplace challenge.
package visitorpatt

// Node represents a tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Accept calls the visitor function on the node's value, then recurses.
func (n *Node) Accept(visitor func(int)) {
	// TODO(candidate): if n == nil return. Call visitor(n.Val). Recurse left, then right.
	panic("not implemented")
}
