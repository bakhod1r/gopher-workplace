// Package baseconv formats a non-negative int in an arbitrary base 2..16.
package baseconv

// Format returns n written in the given base (2..16) using lowercase digits
// 0-9a-f. n is non-negative; Format(0, base) is "0".
//
// TODO(candidate): implement by repeated division, building digits.
func Format(n, base int) string {
	panic("not implemented")
}
