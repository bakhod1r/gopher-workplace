# Reattaching returned subtrees

## Intuition

Recursive tree mutators return the (possibly new) subtree; the parent must reassign it, or newly created nodes are lost.

## Approach

1. `Insert` returns the (possibly new) subtree root.
2. The bug discards that return, so a new leaf is never linked in.
3. Reassign: `root.Left = Insert(root.Left, v)`.

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

Inserting into an empty child, `Insert(root.Left, v)` builds a node but throws it away. Assigning the result back attaches it to `root.Left`.

## Pitfalls

- `Insert(root.Left, v)` alone drops the result for empty subtrees.
- Always write `root.Left = Insert(root.Left, v)`.
