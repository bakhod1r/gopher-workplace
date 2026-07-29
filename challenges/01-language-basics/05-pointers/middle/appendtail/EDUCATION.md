# Updating a head pointer via double indirection

## Intuition

Appending to an empty list must reseat the caller's head; a `**Node` lets the callee do that, while non-empty lists just link at the tail.

## Approach

1. `head` is a `**Node` so the empty-list case can set the caller's head.
2. If `*head == nil`, assign a new node and return.
3. Otherwise walk to the tail and link the new node there.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Append(head **Node, v int) {
	if *head == nil {
		*head = &Node{Val: v}
		return
	}
	n := *head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &Node{Val: v}
}
```

## Walkthrough

On an empty list `Append(&h, 1)` writes the head via `*head`. `Append(&h, 2)` walks to node 1 and sets its `Next`, giving `1->2`.

## Pitfalls

- Without `**Node`, an append to an empty list can't update the caller.
- Walk `n.Next` until nil to find the tail.
