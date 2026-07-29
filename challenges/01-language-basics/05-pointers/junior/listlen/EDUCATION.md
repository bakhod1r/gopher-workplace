# Traversing pointer-linked structures

## Intuition

Following `Next` pointers to nil visits every node; nil marks the end.

## Approach

1. Walk from `head`, following `.Next`.
2. Count each non-nil node.
3. Stop when the pointer is nil.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Length(head *Node) int {
	count := 0
	for n := head; n != nil; n = n.Next {
		count++
	}
	return count
}
```

## Walkthrough

For `1 -> 2 -> 3`: the loop visits three nodes before `n` becomes nil, returning `3`. An empty list starts nil, returning `0`.

## Pitfalls

- Stop at nil, don't dereference it.
- An empty list (nil head) has length 0.
