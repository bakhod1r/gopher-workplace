# Binary Search Tree

## Intuition

Returning the subtree root is the standard trick that lets one function handle both "create the node" and "descend into an existing one" without pointer-to-pointer juggling.

## Approach

1. `Insert`: create a node when the root is nil; otherwise recurse left or right and ignore equals.
2. `InOrder`: append the left subtree, the value, then the right subtree.

## Solution

```go
func Insert[T cmp.Ordered](root *TreeNode[T], v T) *TreeNode[T] {
	if root == nil {
		return &TreeNode[T]{Value: v}
	}
	switch {
	case v < root.Value:
		root.Left = Insert(root.Left, v)
	case v > root.Value:
		root.Right = Insert(root.Right, v)
	}
	return root
}

func InOrder[T cmp.Ordered](root *TreeNode[T]) []T {
	out := make([]T, 0)
	if root == nil {
		return out
	}
	out = append(out, InOrder(root.Left)...)
	out = append(out, root.Value)
	out = append(out, InOrder(root.Right)...)
	return out
}
```

## Walkthrough

Inserting 2, then 1, then 3 puts 1 to the left and 3 to the right; the in-order walk emits `1 2 3`.

## Pitfalls

- Not assigning the result of the recursive `Insert` back into `Left`/`Right`, so new nodes are dropped.
- Inserting duplicates, which the spec says to ignore.
- Walking node-left-right, which is pre-order, not sorted.
