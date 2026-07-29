# Dropping references with nil

## Intuition

Assigning nil to a pointer field removes the link; if no other reference remains, the target becomes eligible for garbage collection.

## Approach

1. `n` is a `*Node`.
2. `n.Next = nil` drops the link; `Value` is left intact.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func Detach(n *Node) {
	n.Next = nil
}
```

## Walkthrough

`Detach(n)` writes `nil` into `n.Next`, so the node no longer references its successor.

## Pitfalls

- Only the Next field changes; Val stays.
- Nil-ing references is how you let the GC reclaim memory.
