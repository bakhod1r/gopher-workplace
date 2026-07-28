// Package rotate does circular bit rotation on a byte.
package rotate

// Left rotates the 8 bits of b left by n positions (n may be >= 8). Bits shifted
// off the top re-enter at the bottom.
//
// TODO(candidate): implement with shifts, masking n to 0..7.
func Left(b byte, n int) byte {
	panic("not implemented")
}
