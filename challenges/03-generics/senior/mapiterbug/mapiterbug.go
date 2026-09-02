// Package mapiterbug — Gopher Workplace challenge.
package mapiterbug

// Mode returns the most frequent element and true.
// On a tie the earliest-occurring element wins.
//
// Examples:
//
//	Mode([]int{1, 1, 2, 2}) => 1, true
func Mode[T comparable](s []T) (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	counts := make(map[T]int, len(s))
	for _, v := range s {
		counts[v]++
	}
	var best T
	bestN := 0
	for v, n := range counts {
		if n > bestN {
			best, bestN = v, n
		}
	}
	return best, true
	// CHANGE CODE ABOVE THIS LINE
}
