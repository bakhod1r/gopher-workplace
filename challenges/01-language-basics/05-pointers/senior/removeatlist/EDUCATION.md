# Reaching the predecessor for deletion

## Intuition

Unlinking a node needs a handle on the node before it; walking all the way to the target leaves nothing to relink.

## Approach

1. To delete index `i` you must stop at index `i-1` (the node before it).
2. The bug walks to index `i`; loop to `i-1` instead.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func RemoveAt(head *Node, i int) *Node {
	if i == 0 {
		return head.Next
	}
	prev := head
	for k := 0; k < i-1; k++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	return head
}
```

## Walkthrough

Walking `k < i` lands `prev` on the target itself, so the splice skips the wrong node. `k < i-1` stops one before, and `prev.Next = prev.Next.Next` removes index `i`.

## Pitfalls

- To delete index i, stop `prev` at i-1.
- Then `prev.Next = prev.Next.Next` skips the target.
