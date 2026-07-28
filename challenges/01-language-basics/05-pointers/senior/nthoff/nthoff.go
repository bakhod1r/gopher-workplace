// Package nthoff returns the n-th node from the end. A planted bug advances the
// lead pointer n+1 steps (or the wrong count), returning the wrong node.
package nthoff

type Node struct {
	Val  int
	Next *Node
}

// NthFromEnd returns the n-th node from the end (1-based). Assume n <= length.
func NthFromEnd(head *Node, n int) *Node {
	lead := head
	// CHANGE CODE BELOW THIS LINE
	for i := 0; i <= n; i++ {
		// CHANGE CODE ABOVE THIS LINE
		lead = lead.Next
	}
	trail := head
	for lead != nil {
		lead = lead.Next
		trail = trail.Next
	}
	return trail
}
