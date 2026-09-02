// Package dropwhilebug — Gopher Workplace challenge.
package dropwhilebug

// DropWhile returns s without its leading elements satisfying pred.
// Later matching elements are kept.
//
// Examples:
//
//	DropWhile([]int{2, 4, 5, 6}, isEven) => []int{5, 6}
func DropWhile[T any](s []T, pred func(T) bool) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, len(s))
	for _, v := range s {
		if pred(v) {
			continue
		}
		out = append(out, v)
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
