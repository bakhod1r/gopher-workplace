// Package delhead removes the first node. A planted bug always returns the
// original head, so deleting the head node doesn't actually drop it.
package delhead

type Node struct {
	Val  int
	Next *Node
}

// RemoveFirst removes the first node and returns the new head.
func RemoveFirst(head *Node) *Node {
	if head == nil {
		return nil
	}
	// CHANGE CODE BELOW THIS LINE
	return head
	// CHANGE CODE ABOVE THIS LINE
}
