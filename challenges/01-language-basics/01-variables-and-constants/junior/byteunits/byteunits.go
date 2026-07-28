// Package byteunits — Gopher Workplace challenge.
package byteunits

// TODO(candidate): declare the binary storage constants KiB, MiB, GiB.
// Each step up is 1024x (2^10) the previous one:
//
//	KiB = 1024
//	MiB = 1024 * 1024
//	GiB = 1024 * 1024 * 1024
//
// Derive them from iota (do not hand-type each number).
const (
	KiB = 0
	MiB = 0
	GiB = 0
)

// Bytes returns n whole KiB expressed in bytes.
func Bytes(n int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
