// Package sumtreebug sums a tree. A planted bug returns only the root value,
// never recursing into the children.
package sumtreebug

type Tree struct {
	Val         int
	Left, Right *Tree
}

// SumTree returns the sum of all node values.
func SumTree(t *Tree) int {
	if t == nil {
		return 0
	}
	// CHANGE CODE BELOW THIS LINE
	return t.Val
	// CHANGE CODE ABOVE THIS LINE
}
