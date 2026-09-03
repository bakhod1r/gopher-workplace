// Package escapecallback — Gopher Workplace challenge.
package escapecallback

// Each calls f for every element of s.
//
// It is a package-level variable, so the compiler cannot see which function
// runs and must assume the callback it is given escapes.
var Each = func(s []int, f func(int)) {
	for _, v := range s {
		f(v)
	}
}

// Sum returns the total of s.
//
// The helper Each is not inlinable, so any closure handed to it escapes —
// and with it everything the closure captures.
//
// Examples:
//
//	Sum([]int{1, 2, 3}) => 6
func Sum(s []int) int64 {
	// CHANGE CODE BELOW THIS LINE
	var total int64
	Each(s, func(v int) { total += int64(v) })
	return total
	// CHANGE CODE ABOVE THIS LINE
}
