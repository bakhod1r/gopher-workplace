// Package graphgen — Gopher Workplace challenge.
package graphgen

// Graph is a directed graph over nodes of K.
// Use NewGraph to create one.
type Graph[K comparable] struct {
	edges map[K][]K
}

// NewGraph returns an empty directed graph.
func NewGraph[K comparable]() *Graph[K] {
	// TODO(candidate): allocate the adjacency map.
	panic("not implemented")
}

// AddEdge records a directed edge from a to b.
// Duplicate edges are ignored.
func (g *Graph[K]) AddEdge(a, b K) {
	// TODO(candidate): record the edge, skipping duplicates.
	panic("not implemented")
}

// Neighbors returns the nodes reachable from a in one step,
// in insertion order.
func (g *Graph[K]) Neighbors(a K) []K {
	// TODO(candidate): return a copy of the adjacency list.
	panic("not implemented")
}

// Degree returns how many edges leave a.
func (g *Graph[K]) Degree(a K) int {
	// TODO(candidate): report the out-degree.
	panic("not implemented")
}
