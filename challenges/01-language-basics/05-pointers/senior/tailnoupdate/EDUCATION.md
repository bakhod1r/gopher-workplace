# Growing an empty list

## Intuition

With a `*Node` signature the empty case must return the new head; returning nil loses the value.

## Approach

1. Appending to an empty list must create the first node.
2. The bug returns nil, dropping the value.
3. Return `&Node{Val: v}` in the empty case.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Append(head *Node, v int) *Node {
	if head == nil {
		return &Node{Val: v}
	}
	n := head
	for n.Next != nil {
		n = n.Next
	}
	n.Next = &Node{Val: v}
	return head
}
```

## Walkthrough

`Append(nil, 1)` should yield a one-node list; returning nil loses the value. Allocating a node fixes the base case.

## Pitfalls

- An empty list has no tail — build the head instead.
- Return it so the caller can adopt the new head.
