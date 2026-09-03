# One Allocation, Not One Per Element

## Intuition

A pointer needs an address, not its own allocation. Carving one array into n addressable elements gives the same API with two allocations, and the collector sees one object instead of n.

## Approach

1. Allocate `block := make([]Node, n)` and `out := make([]*Node, n)`.
2. Set each node's ID and point `out[i]` at `&block[i]`.

## Solution

```go
// Node is one element of the built collection.
type Node struct {
	ID   int
	Next *Node
}

// Build returns n nodes, each pointed at by one element of the result.
//
// Allocating each node separately costs n allocations and scatters them
// across the heap; one backing array plus n pointers into it costs two.
//
// Examples:
//
// 	Build(3) => three nodes with IDs 0, 1, 2
func Build(n int) []*Node {
	if n <= 0 {
		return nil
	}
	block := make([]Node, n)
	out := make([]*Node, n)
	for i := 0; i < n; i++ {
		block[i].ID = i
		out[i] = &block[i]
	}
	return out
}
```

## Walkthrough

Building 256 nodes costs 256 allocations before the fix and 2 after. The trade-off is that the block is freed only when the last pointer into it is gone.

## Pitfalls

- Taking `&node` of a loop variable and expecting distinct nodes — every iteration must address a distinct element.
- Applying this when nodes have wildly different lifetimes; the block is all-or-nothing.
