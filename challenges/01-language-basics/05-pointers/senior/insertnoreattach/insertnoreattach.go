// Package insertnoreattach inserts into a BST. A planted bug calls Insert on a
// child but discards the returned subtree, so inserts into empty subtrees are
// lost.
package insertnoreattach

type Tree struct {
	Val         int
	Left, Right *Tree
}

// Insert adds v into the BST and returns the (possibly new) root.
func Insert(root *Tree, v int) *Tree {
	if root == nil {
		return &Tree{Val: v}
	}
	if v < root.Val {
		// CHANGE CODE BELOW THIS LINE
		Insert(root.Left, v)
		// CHANGE CODE ABOVE THIS LINE
	} else {
		root.Right = Insert(root.Right, v)
	}
	return root
}
