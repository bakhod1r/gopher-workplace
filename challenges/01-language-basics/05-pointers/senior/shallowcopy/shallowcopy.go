// Package shallowcopy copies a tree. A planted bug copies only the root node and
// reuses the original children, so mutating the copy's subtree corrupts the
// original.
package shallowcopy

type Tree struct {
	Val         int
	Left, Right *Tree
}

// Copy returns a deep, independent copy of t.
func Copy(t *Tree) *Tree {
	if t == nil {
		return nil
	}
	// CHANGE CODE BELOW THIS LINE
	return &Tree{Val: t.Val, Left: t.Left, Right: t.Right}
	// CHANGE CODE ABOVE THIS LINE
}
