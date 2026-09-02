// Package tildeconvbug — Gopher Workplace challenge.
package tildeconvbug

// Total sums values of any named type whose underlying type is int64.
// The sum is computed at the full width of the underlying type.
//
// Examples:
//
//	Total([]Millis{3000000000, 3000000000}) => 6000000000
func Total[T ~int64](vals []T) T {
	// CHANGE CODE BELOW THIS LINE
	var sum int32
	for _, v := range vals {
		sum += int32(v)
	}
	return T(sum)
	// CHANGE CODE ABOVE THIS LINE
}
