# Tracking the new head after reversal

## Intuition

The loop's final `prev` is the reversed list's head; the original head has become the tail with a nil Next.

## Approach

1. After the loop, `prev` is the new head; `head` is now the tail.
2. The bug returns `head` (old head, now last). Return `prev`.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Reverse(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
```

## Walkthrough

The reversal is correct but returning `head` hands back the original first node — now the tail. `prev` is the new front, `3->2->1`.

## Pitfalls

- After reversal, `head.Next == nil` (it's the tail).
- Return `prev`, the last node advanced.
