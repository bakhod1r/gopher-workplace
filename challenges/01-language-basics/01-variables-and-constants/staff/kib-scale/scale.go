// Package scale converts kibibytes to bytes. Planted decimal factor.
package scale

// KiB is the binary kilobyte factor.
// CHANGE CODE BELOW THIS LINE
const KiB = 1000

// CHANGE CODE ABOVE THIS LINE

// Bytes returns n kibibytes in bytes.
func Bytes(n int) int { return n * KiB }
