# In-place tree transformation

## Intuition

Swapping children at every node and recursing mirrors the tree; the swap is a parallel assignment on pointer fields.

## Approach

1. Nil → return.
2. Swap `t.Left` and `t.Right`.
3. Recurse into both subtrees.

## Solution

```go
type Tree struct {
	Val         int
	Left, Right *Tree
}

func Mirror(t *Tree) {
	if t == nil {
		return
	}
	t.Left, t.Right = t.Right, t.Left
	Mirror(t.Left)
	Mirror(t.Right)
}
```

## Walkthrough

`Mirror` swaps each node's children top-down; a root with children 2,3 ends with children 3,2, and the swap propagates to every subtree.

## Pitfalls

- Swap before or after recursing — both work since the swap is local.
- Base-case nil returns immediately.
