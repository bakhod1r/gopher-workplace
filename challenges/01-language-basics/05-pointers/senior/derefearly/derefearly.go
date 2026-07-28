// Package derefearly reads a node's value safely. A planted bug reads head.Val
// before checking head for nil, panicking on an empty list.
package derefearly

type Node struct {
	Val  int
	Next *Node
}

// FirstOr returns head.Val, or def when head is nil.
func FirstOr(head *Node, def int) int {
	// CHANGE CODE BELOW THIS LINE
	v := head.Val
	if head == nil {
		return def
	}
	return v
	// CHANGE CODE ABOVE THIS LINE
}
