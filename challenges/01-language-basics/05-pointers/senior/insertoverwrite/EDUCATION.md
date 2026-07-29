# Non-destructive BST insertion

## Intuition

Insertion must recurse to a nil slot and add a node there; overwriting the root each call collapses the tree to one node.

## Approach

1. The bug replaces the whole tree with a new root on every call.
2. Recurse into the correct subtree and only allocate on nil: `if root == nil { return &Tree{Val: v} }`.

## Solution

```go
type Tree struct {
	Val         int
	Left, Right *Tree
}

func Insert(root *Tree, v int) *Tree {
	if root == nil {
		return &Tree{Val: v}
	}
	if v < root.Val {
		root.Left = Insert(root.Left, v)
	} else {
		root.Right = Insert(root.Right, v)
	}
	return root
}
```

## Walkthrough

`root = &Tree{Val: v}` discards the existing tree, so only the last insert survives. Recursing left/right by BST order preserves prior nodes.

## Pitfalls

- `root = &Tree{Val: v}` discards everything already inserted.
- Recurse to the correct empty child instead.
