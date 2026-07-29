# Building a binary search tree

## Intuition

Insertion recurses to the correct nil slot and returns each subtree so parents can relink; the root return covers the empty tree.

## Approach

1. Nil subtree → return a new node.
2. Smaller values recurse left, larger recurse right.
3. Reassign the child link and return `root`.

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

`Insert(root(5), 3)`: `3 < 5`, so recurse left into nil, which returns `&Tree{Val:3}`, attached as `root.Left`.

## Pitfalls

- Returning the subtree lets the parent reattach it.
- The nil slot is where the new node goes.
