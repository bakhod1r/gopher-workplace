// Package bfsmarkbug — Gopher Workplace challenge.
package bfsmarkbug

// BFS returns the nodes reachable from start in breadth-first order.
// Each node appears exactly once.
func BFS[T comparable](adj map[T][]T, start T) []T {
	// CHANGE CODE BELOW THIS LINE
	seen := make(map[T]bool)
	queue := []T{start}
	out := make([]T, 0)
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		seen[n] = true
		out = append(out, n)
		for _, m := range adj[n] {
			if !seen[m] {
				queue = append(queue, m)
			}
		}
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
