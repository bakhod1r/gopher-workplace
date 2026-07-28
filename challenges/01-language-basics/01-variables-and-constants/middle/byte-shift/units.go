// Package units defines binary byte-size constants with iota shifting.
package units

// ByteSize counts bytes.
type ByteSize uint64

// KB, MB, GB, TB as 1<<(10*n) using iota.
//
// TODO(candidate): define with a single iota expression 1 << (10 * iota),
// skipping the iota==0 slot with the blank identifier.
const (
	_  ByteSize = 0
	KB ByteSize = 0
	MB ByteSize = 0
	GB ByteSize = 0
	TB ByteSize = 0
)

// Humanize returns the largest unit that divides n evenly, as (count, symbol).
// For n < KB it returns (n, "B"). Prefer larger units on ties.
//
// TODO(candidate): implement.
func Humanize(n ByteSize) (ByteSize, string) {
	panic("not implemented")
}
