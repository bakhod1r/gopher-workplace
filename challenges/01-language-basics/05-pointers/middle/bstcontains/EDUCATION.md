# Searching a BST

## Intuition

The ordering invariant lets search discard one subtree per step, giving O(height) lookups.

## Approach

1. Nil → not found.
2. Equal → found.
3. Otherwise recurse into the side dictated by BST order.

## Solution

```go
type Tree struct {
	Val         int
	Left, Right *Tree
}

func Contains(root *Tree, v int) bool {
	if root == nil {
		return false
	}
	if v == root.Val {
		return true
	}
	if v < root.Val {
		return Contains(root.Left, v)
	}
	return Contains(root.Right, v)
}
```

## Walkthrough

`Contains(bst, 8)`: at each node compare and branch; the search reaches node 8 and returns true. A value not present bottoms out at nil → false.

## Pitfalls

- Base-case nil returns false.
- Compare then recurse into exactly one child.
