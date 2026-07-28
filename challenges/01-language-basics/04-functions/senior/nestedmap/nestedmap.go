// Package nestedmap records edges in an adjacency map. A planted bug writes to
// an inner map that was never created, panicking on the first insert.
package nestedmap

// Add records a directed edge from -> to in the adjacency structure.
func Add(g map[string]map[string]bool, from, to string) {
	// CHANGE CODE BELOW THIS LINE
	g[from][to] = true
	// CHANGE CODE ABOVE THIS LINE
}
