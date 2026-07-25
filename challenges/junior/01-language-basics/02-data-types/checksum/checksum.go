// Package checksum — Gopher Workplace challenge.
package checksum

// Checksum returns the modulo-256 sum of every byte in data: the bytes added
// together as uint8, letting the fixed 8-bit width wrap around on overflow.
// A nil or empty slice checksums to 0.
//
// Examples:
//
//	Checksum([]byte{1, 2, 3})      => 6
//	Checksum([]byte{255, 1})       => 0   // 256 wraps to 0
//	Checksum([]byte{200, 100})     => 44  // 300 mod 256
func Checksum(data []byte) uint8 {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
