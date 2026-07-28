// Package parity computes the parity bit of a value.
// A planted bug sums bits instead of XOR-ing them.
package parity

// Parity returns 1 if x has an odd number of set bits, else 0.
func Parity(x uint32) int {
	p := 0
	for x != 0 {
		// CHANGE CODE BELOW THIS LINE
		p += int(x & 1)
		// CHANGE CODE ABOVE THIS LINE
		x >>= 1
	}
	return p
}
