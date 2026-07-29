# Handling leftovers in a merge

## Intuition

The merge loop stops when either list empties; the non-empty remainder must be linked to the tail to avoid dropping nodes.

## Approach

1. After the main loop one list still has nodes.
2. The bug drops that remainder; attach whichever of `a`/`b` is non-nil to `tail.Next`.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Merge(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}
```

## Walkthrough

When `a` runs out with `3->4->5` left in `b`, the tail must link to `b`. Omitting this truncates the merged list.

## Pitfalls

- The loop condition `a != nil && b != nil` leaves one list non-empty.
- Link the leftover directly; it's already sorted.
