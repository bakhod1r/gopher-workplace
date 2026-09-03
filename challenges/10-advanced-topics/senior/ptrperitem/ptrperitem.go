// Package ptrperitem — Gopher Workplace challenge.
package ptrperitem

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
//	Build(3) => three nodes with IDs 0, 1, 2
func Build(n int) []*Node {
	// CHANGE CODE BELOW THIS LINE
	if n <= 0 {
		return nil
	}
	out := make([]*Node, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, &Node{ID: i})
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
