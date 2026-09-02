// Package treegen — Gopher Workplace challenge.
package treegen

import (
	"cmp"
)

// TreeNode is one node of a binary search tree of T.
type TreeNode[T cmp.Ordered] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Insert adds v to the tree, ignoring duplicates.
func Insert[T cmp.Ordered](root *TreeNode[T], v T) *TreeNode[T] {
	// TODO(candidate): insert v, returning the (possibly new) root.
	panic("not implemented")
}

// InOrder returns the values in ascending order.
func InOrder[T cmp.Ordered](root *TreeNode[T]) []T {
	// TODO(candidate): walk left, node, right.
	panic("not implemented")
}
