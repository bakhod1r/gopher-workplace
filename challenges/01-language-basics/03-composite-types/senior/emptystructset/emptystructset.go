// Package emptystructset builds a set. A planted bug tests membership by the
// stored bool value instead of presence, but stores false.
package emptystructset

// Intersect returns values present in both a and b (unique, order of a).
func Intersect(a, b []int) []int {
	inB := make(map[int]bool)
	for _, x := range b {
		// CHANGE CODE BELOW THIS LINE
		inB[x] = false
		// CHANGE CODE ABOVE THIS LINE
	}
	out := []int{}
	seen := make(map[int]bool)
	for _, x := range a {
		if inB[x] && !seen[x] {
			out = append(out, x)
			seen[x] = true
		}
	}
	return out
}
