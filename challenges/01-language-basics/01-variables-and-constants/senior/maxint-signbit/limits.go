// Package limits derives the max signed int from a bit pattern.
// A planted bug forgets the sign-bit shift.
package limits

// allBits is every bit set; kept as a var so the conversion below happens at
// run time (a constant conversion would overflow int at compile time).
var allBits = ^uint(0)

// MaxInt is meant to be the largest positive int.
// CHANGE CODE BELOW THIS LINE
var MaxInt = int(allBits)

// CHANGE CODE ABOVE THIS LINE

// Overflows reports whether adding 1 to MaxInt would overflow (it always does).
func Overflows() bool { return MaxInt+1 < MaxInt }
