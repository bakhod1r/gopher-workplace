// Package insertoverwrite inserts into a BST. A planted bug reassigns root to the
// new node on every call, discarding the existing tree.
package insertoverwrite

type Tree struct {
	Val         int
	Left, Right *Tree
}

// Insert adds v into the BST and returns the root.
func Insert(root *Tree, v int) *Tree {
	// CHANGE CODE BELOW THIS LINE
	root = &Tree{Val: v}
	return root
	// CHANGE CODE ABOVE THIS LINE
}
