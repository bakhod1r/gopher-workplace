// Package filterinplacebug — Gopher Workplace challenge.
package filterinplacebug

// FilterInPlace keeps the elements satisfying pred, reusing s's storage.
// It returns the shortened slice.
//
// Examples:
//
//	FilterInPlace([]int{1, 2, 3}, even) => []int{2}
func FilterInPlace[T any](s []T, pred func(T) bool) []T {
	// CHANGE CODE BELOW THIS LINE
	n := 0
	for _, v := range s {
		if pred(v) {
			s[n] = v
			n++
		}
	}
	return s
	// CHANGE CODE ABOVE THIS LINE
}
