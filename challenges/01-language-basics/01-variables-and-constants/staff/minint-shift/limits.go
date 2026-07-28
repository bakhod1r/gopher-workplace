// Package limits builds MinInt from a signed shift. Planted off-by-one shift.
package limits

// MinInt is meant to be the most negative int64-width value.
// CHANGE CODE BELOW THIS LINE
const MinInt = -1 << 62

// CHANGE CODE ABOVE THIS LINE

// SymmetricTo reports whether -MinInt would overflow (it must: no positive
// counterpart exists in two's complement).
func SymmetricTo() bool { return MinInt < -(MinInt + 1) }
