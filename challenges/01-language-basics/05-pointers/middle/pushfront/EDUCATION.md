# Prepending to a linked list

## Intuition

A new node pointing at the current head becomes the new head — an O(1) push front.

## Approach

1. Allocate a node whose `Next` is the current `head`.
2. Return it as the new head.
3. The old list is reused untouched.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func PushFront(head *Node, v int) *Node {
	return &Node{Val: v, Next: head}
}
```

## Walkthrough

`PushFront(1->2, 0)` builds `&Node{Val:0, Next: head}`; that node links to `1->2` and is returned, giving `0->1->2`.

## Pitfalls

- Return the new node; the old head is now second.
- Works with a nil head (empty list).
