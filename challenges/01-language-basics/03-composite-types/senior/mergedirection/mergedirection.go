// Package mergedirection merges override config over base. A planted bug copies
// in the wrong order, so base wins instead of the override.
package mergedirection

// Merge returns base with over applied on top (over wins on collisions).
func Merge(base, over map[string]int) map[string]int {
	out := make(map[string]int)
	// CHANGE CODE BELOW THIS LINE
	for k, v := range over {
		out[k] = v
	}
	for k, v := range base {
		out[k] = v
	}
	// CHANGE CODE ABOVE THIS LINE
	return out
}
