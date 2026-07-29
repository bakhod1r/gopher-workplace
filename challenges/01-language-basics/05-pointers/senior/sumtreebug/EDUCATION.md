# Recursing into all children

## Intuition

A tree aggregation must combine the node with BOTH subtree results; returning only the node value discards the rest.

## Approach

1. The bug returns only `t.Val`, ignoring the subtrees.
2. Return `t.Val + SumTree(t.Left) + SumTree(t.Right)`.

## Solution

```go
type Tree struct {
	Val         int
	Left, Right *Tree
}

func SumTree(t *Tree) int {
	if t == nil {
		return 0
	}
	return t.Val + SumTree(t.Left) + SumTree(t.Right)
}
```

## Walkthrough

Returning just `t.Val` sums only the root. Adding both recursive subtree sums accumulates the whole tree to 10.

## Pitfalls

- `return t.Val` visits one node only.
- Add `SumTree(t.Left) + SumTree(t.Right)`.
