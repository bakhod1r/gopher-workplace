// Package nestedmapinit counts pairs in a nested map. A planted bug writes to an
// uninitialized inner map.
package nestedmapinit

// Tally counts occurrences of (outer, inner) pairs into a nested map.
func Tally(pairs [][2]string) map[string]map[string]int {
	m := make(map[string]map[string]int)
	for _, p := range pairs {
		o, i := p[0], p[1]
		// CHANGE CODE BELOW THIS LINE
		// (inner map may not exist yet)
		// CHANGE CODE ABOVE THIS LINE
		m[o][i]++
	}
	return m
}
