// Package setopsmutatebug — Gopher Workplace challenge.
package setopsmutatebug

// Union returns a new set holding every element of a and b.
// Neither input is modified.
//
// Examples:
//
//	Union({1,2}, {2,3}) => {1,2,3}
func Union[T comparable](a, b map[T]struct{}) map[T]struct{} {
	// CHANGE CODE BELOW THIS LINE
	out := a
	for k := range b {
		out[k] = struct{}{}
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
