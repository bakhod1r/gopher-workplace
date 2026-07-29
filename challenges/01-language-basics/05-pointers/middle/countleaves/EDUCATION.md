# Classifying nodes during traversal

## Intuition

A leaf test (both children nil) plus a recursive sum counts leaves; the same shape counts internal nodes or full nodes.

## Approach

1. Nil → 0.
2. A node with no children is a leaf → 1.
3. Otherwise sum the leaves of both subtrees.

## Solution

```go
type Tree struct {
	Val         int
	Left, Right *Tree
}

func CountLeaves(t *Tree) int {
	if t == nil {
		return 0
	}
	if t.Left == nil && t.Right == nil {
		return 1
	}
	return CountLeaves(t.Left) + CountLeaves(t.Right)
}
```

## Walkthrough

For a root with two childless children, each child returns 1, so the root returns `1 + 1 = 2`.

## Pitfalls

- Distinguish nil (0) from a leaf (1).
- Sum both subtrees for internal nodes.
