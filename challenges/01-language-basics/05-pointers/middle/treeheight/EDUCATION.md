# Recursion over binary trees

## Intuition

Tree algorithms recurse on Left and Right, base-casing on nil; height combines the deeper subtree.

## Approach

1. Base: nil has height 0.
2. Recurse into both children.
3. Return `1 + max(left, right)`.

## Solution

```go
type Tree struct {
	Val         int
	Left, Right *Tree
}

func Height(t *Tree) int {
	if t == nil {
		return 0
	}
	l, r := Height(t.Left), Height(t.Right)
	if l > r {
		return 1 + l
	}
	return 1 + r
}
```

## Walkthrough

For a balanced 3-level tree each side returns 2, so the root returns `1 + 2 = 3`.

## Pitfalls

- Base-case nil before dereferencing.
- Height uses max of subtrees; size would use sum.
