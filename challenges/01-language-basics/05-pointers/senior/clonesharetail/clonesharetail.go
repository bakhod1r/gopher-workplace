// Package clonesharetail copies a list. A planted bug copies only the head node
// and reuses the original tail, so mutating the copy's tail corrupts the source.
package clonesharetail

type Node struct {
	Val  int
	Next *Node
}

// Copy returns a deep copy of the list.
func Copy(head *Node) *Node {
	if head == nil {
		return nil
	}
	// CHANGE CODE BELOW THIS LINE
	return &Node{Val: head.Val, Next: head.Next}
	// CHANGE CODE ABOVE THIS LINE
}
