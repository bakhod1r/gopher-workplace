# Deep-copying tree structures

## Intuition

Copying only the root while reusing child pointers aliases the subtrees; a true deep copy recurses on both children.

## Approach

1. Copying `Left`/`Right` by reference shares subtrees with the original.
2. Recurse: `Left: Copy(t.Left), Right: Copy(t.Right)`.

## Solution

```go
type Tree struct {
	Val         int
	Left, Right *Tree
}

func Copy(t *Tree) *Tree {
	if t == nil {
		return nil
	}
	return &Tree{Val: t.Val, Left: Copy(t.Left), Right: Copy(t.Right)}
}
```

## Walkthrough

The bug reuses the original child pointers, so mutating the copy mutates the source. Recursively copying children makes the trees independent.

## Pitfalls

- `Left: t.Left` shares the original subtree.
- Recurse with `Copy(t.Left)` for independence.
