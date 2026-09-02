# Adjacency Graph

## Intuition

Storing neighbours in insertion order rather than a set keeps traversals deterministic, at the cost of the linear duplicate check.

## Approach

1. `AddEdge`: return early when the edge exists, otherwise append.
2. `Neighbors`: copy the stored slice.
3. `Degree`: return its length.

## Solution

```go
func NewGraph[K comparable]() *Graph[K] {
	return &Graph[K]{edges: make(map[K][]K)}
}

func (g *Graph[K]) AddEdge(a, b K) {
	for _, n := range g.edges[a] {
		if n == b {
			return
		}
	}
	g.edges[a] = append(g.edges[a], b)
}

func (g *Graph[K]) Neighbors(a K) []K {
	out := make([]K, len(g.edges[a]))
	copy(out, g.edges[a])
	return out
}

func (g *Graph[K]) Degree(a K) int {
	return len(g.edges[a])
}
```

## Walkthrough

`AddEdge(a,b)` twice finds `b` already present the second time and returns without appending.

## Pitfalls

- Returning `g.edges[a]` directly, letting callers rewrite the graph.
- Skipping the duplicate check and inflating the degree.
- Panicking on an unknown node instead of reporting nothing.
