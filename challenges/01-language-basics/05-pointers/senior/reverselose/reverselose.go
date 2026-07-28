// Package reverselose reverses a list. A planted bug overwrites cur.Next before
// saving it, cutting off the rest of the list after one node.
package reverselose

type Node struct {
	Val  int
	Next *Node
}

// Reverse reverses the list and returns the new head.
func Reverse(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		// CHANGE CODE BELOW THIS LINE
		cur.Next = prev
		next := cur.Next
		// CHANGE CODE ABOVE THIS LINE
		prev = cur
		cur = next
	}
	return prev
}
