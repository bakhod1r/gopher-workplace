# Aggregating over a tree

## Intuition

A post-order recursion combines the current value with both subtree sums.

## Approach

1. Nil → 0.
2. Return `t.Val + SumTree(left) + SumTree(right)`.

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

For root 1 with children 2 and 3: `1 + 2 + 3 = 6`.

## Pitfalls

- Base-case nil returns the identity (0 for sums).
- Recurse both children.
