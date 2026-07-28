// Package fletcher16 computes the Fletcher-16 checksum.
// A planted bug uses the wrong modulus.
package fletcher16

// Checksum returns the Fletcher-16 checksum of data: two running sums mod 255,
// combined as (sum2 << 8) | sum1.
func Checksum(data []byte) uint16 {
	var sum1, sum2 uint16
	for _, b := range data {
		// CHANGE CODE BELOW THIS LINE
		sum1 = (sum1 + uint16(b)) % 256
		sum2 = (sum2 + sum1) % 256
		// CHANGE CODE ABOVE THIS LINE
	}
	return sum2<<8 | sum1
}
