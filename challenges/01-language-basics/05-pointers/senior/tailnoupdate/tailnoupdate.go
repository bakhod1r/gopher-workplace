// Package tailnoupdate appends to a list's tail. A planted bug takes *Node (not
// **Node) and returns nothing, so appending to an empty list can't tell the
// caller about the new head. It is fixed by returning the head.
package tailnoupdate

type Node struct {
	Val  int
	Next *Node
}

// Append adds v at the tail and returns the (possibly new) head.
func Append(head *Node, v int) *Node {
	if head == nil {
		// CHANGE CODE BELOW THIS LINE
		return nil
		// CHANGE CODE ABOVE THIS LINE
	}
	n := head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &Node{Val: v}
	return head
}
