// Package removeatlist removes the node at index i (0-based). A planted bug walks
// to the target node itself instead of its predecessor, so it can't relink.
package removeatlist

type Node struct {
	Val  int
	Next *Node
}

// RemoveAt removes the node at index i (0 <= i < length) and returns the head.
func RemoveAt(head *Node, i int) *Node {
	if i == 0 {
		return head.Next
	}
	prev := head
	// CHANGE CODE BELOW THIS LINE
	for k := 0; k < i; k++ {
		prev = prev.Next
	}
	// CHANGE CODE ABOVE THIS LINE
	prev.Next = prev.Next.Next
	return head
}
