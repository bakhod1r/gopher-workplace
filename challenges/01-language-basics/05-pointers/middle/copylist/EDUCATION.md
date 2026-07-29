# Deep vs shallow copy

## Intuition

Copying the head pointer shares the whole list; a deep copy allocates a new node per element so mutations don't leak.

## Approach

1. Base case: nil copies to nil.
2. Allocate a new node with the same `Val` and recursively copy the tail.
3. No node is shared with the original.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Copy(head *Node) *Node {
	if head == nil {
		return nil
	}
	return &Node{Val: head.Val, Next: Copy(head.Next)}
}
```

## Walkthrough

`Copy(1->2->3)` builds a new node 1 whose Next is a fresh copy of `2->3`, producing a fully independent list.

## Pitfalls

- Returning `head` (shallow) shares every node.
- Allocate a new node for each; recurse on Next.
