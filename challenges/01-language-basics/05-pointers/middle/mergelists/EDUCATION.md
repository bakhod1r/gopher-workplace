# Merging with a dummy head

## Intuition

A sentinel dummy node removes special-casing the first append; a tail pointer builds the merged list in order.

## Approach

1. Use a `dummy` head and a `tail` cursor.
2. While both lists are non-empty, attach the smaller node and advance it.
3. Append whatever remains; return `dummy.Next`.

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

`Merge(1->3->5, 2->4->6)`: pick 1, then 2, then 3... alternating by value, building the sorted `1->2->3->4->5->6`.

## Pitfalls

- The dummy head avoids branching on the first node.
- Attach the leftover list directly once one runs out.
