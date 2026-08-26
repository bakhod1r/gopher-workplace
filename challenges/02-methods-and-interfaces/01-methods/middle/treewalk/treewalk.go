// Package treewalk — Gopher Workplace challenge.
package treewalk

// Tree is a binary tree node.
type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

// Walk returns an in-order traversal of the tree as a slice.
// For a nil tree, return an empty (non-nil) slice.
//
// Examples:
//
//	tree(2, tree(1), tree(3)).Walk() => [1, 2, 3]
//	(*Tree)(nil).Walk()               => []
func (t *Tree) Walk() []int {
	// TODO(candidate): in-order traversal — left, root, right.
	panic("not implemented")
}
