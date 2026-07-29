# Statement ordering in pointer surgery

## Intuition

Overwriting `cur.Next` before saving it loses the rest of the list; capture the successor first.

## Approach

1. You must save `next := cur.Next` **before** overwriting `cur.Next`.
2. The bug reverses the order, so after `cur.Next = prev` the saved `next` is already `prev` — the rest of the list is lost.

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

With the lines swapped, `cur.Next` is set to `prev` and only then read into `next`, so the traversal loops back instead of advancing. Saving `next` first keeps the walk moving forward.

## Pitfalls

- Read `next := cur.Next` before `cur.Next = prev`.
- Reversed order truncates the list after the first node.
