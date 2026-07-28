// Package mapniljson returns per-key counts that must JSON-encode to {} when
// empty. A planted bug returns a nil map, which encodes to null.
package mapniljson

// Counts returns the occurrence count of each string. For empty input it must
// return a non-nil (empty) map so JSON is {} and not null.
func Counts(xs []string) map[string]int {
	// CHANGE CODE BELOW THIS LINE
	var m map[string]int
	// CHANGE CODE ABOVE THIS LINE
	for _, x := range xs {
		if m == nil {
			m = map[string]int{}
		}
		m[x]++
	}
	return m
}
