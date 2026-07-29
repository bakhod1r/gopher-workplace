# Deep-copying a linked list

## Intuition

Copying only the head while reusing Next aliases the whole tail; recurse to allocate a fresh node per element.

## Approach

1. `Next: head.Next` shares the rest of the list with the original.
2. Recurse: `Next: Copy(head.Next)`.

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

The bug copies only the head and reuses the tail, so the lists alias from node 2 on. Recursively copying `Next` makes the copy independent.

## Pitfalls

- `Next: head.Next` shares the rest of the list.
- Use `Next: Copy(head.Next)` for a deep copy.
