# Clearing struct fields through a pointer

## Intuition

A struct pointer reaches every field, including nested pointer fields, letting a helper reset an object in place.

## Approach

1. `n` is a `*Node`.
2. Zero both fields through it: `n.Value = 0; n.Next = nil`.

## Solution

```go
type Node struct {
	Value int
	Next  *Node
}

func Reset(n *Node) {
	n.Value = 0
	n.Next = nil
}
```

## Walkthrough

`Reset(n)` writes `0` into `n.Value` and `nil` into `n.Next`, clearing the caller's node.

## Pitfalls

- Setting `n.Next = nil` drops the reference (helps GC).
- Value receiver would clear a copy only.
