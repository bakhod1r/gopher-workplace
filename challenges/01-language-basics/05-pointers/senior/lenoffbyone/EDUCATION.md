# Initialising accumulators

## Intuition

The loop counts every node, so the counter must start at 0; a start of 1 double-counts and breaks the empty case.

## Approach

1. The counter must start at 0; the bug seeds it at 1.
2. `count := 0` then increment per node.

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

Seeding `count := 1` overcounts by one and returns 1 for an empty list. Starting at 0 gives the true length.

## Pitfalls

- `count := 1` makes an empty list report length 1.
- Let the loop do all the counting from 0.
