// Package clone2dsharebug — Gopher Workplace challenge.
package clone2dsharebug

// Clone2D returns a deep copy of m.
// No row of the result shares storage with a row of m.
//
// Examples:
//
//	Clone2D([][]int{{1}, {2}}) => an independent [][]int{{1}, {2}}
func Clone2D[T any](m [][]T) [][]T {
	// CHANGE CODE BELOW THIS LINE
	out := make([][]T, len(m))
	for i, row := range m {
		out[i] = row
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
