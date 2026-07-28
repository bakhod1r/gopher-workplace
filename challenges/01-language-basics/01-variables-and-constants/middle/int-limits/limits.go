// Package limits derives integer bounds from bit patterns at compile time.
package limits

// MaxUint is the largest uint: all bits set.
// MaxInt is the largest int: all bits set except the sign bit.
// MinInt is the smallest int.
//
// TODO(candidate): define the three using ^uint(0) and shifts, no math package.
const (
	MaxUint = uint(0)
	MaxInt  = int(0)
	MinInt  = int(0)
)

// FitsInInt reports whether the unsigned value v is <= MaxInt.
//
// TODO(candidate): implement.
func FitsInInt(v uint) bool {
	panic("not implemented")
}
