// Package revreturnhead reverses a list. A planted bug returns the original head
// instead of the new head (prev), so the caller only sees the (now last) node.
package revreturnhead

type Node struct {
	Val  int
	Next *Node
}

// Reverse reverses the list and returns the new head.
func Reverse(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// CHANGE CODE BELOW THIS LINE
	return head
	// CHANGE CODE ABOVE THIS LINE
}
