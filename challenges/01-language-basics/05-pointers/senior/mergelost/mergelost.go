// Package mergelost merges two sorted lists. A planted bug forgets to attach the
// leftover of the non-empty list after one runs out, truncating the result.
package mergelost

type Node struct {
	Val  int
	Next *Node
}

// Merge merges two sorted lists and returns the head.
func Merge(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	// CHANGE CODE BELOW THIS LINE

	// CHANGE CODE ABOVE THIS LINE
	return dummy.Next
}
