// Package narrowing — Gopher Workplace challenge.
package narrowing

// ToInt32 converts a 64-bit integer to a 32-bit one. Go requires an explicit
// conversion between numeric types; when the value does not fit in int32, the
// high bits are dropped and the result wraps around (two's-complement), it is
// not clamped.
//
// Examples:
//
//	ToInt32(42)          => 42
//	ToInt32(-7)          => -7
//	ToInt32(2147483647)  => 2147483647   // int32 max, still fits
//	ToInt32(2147483648)  => -2147483648  // one past max, wraps to int32 min
//	ToInt32(4294967296)  => 0            // 2^32 wraps to 0
func ToInt32(n int64) int32 {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
