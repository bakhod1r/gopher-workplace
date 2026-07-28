// Package units builds binary size constants. A planted shift bug breaks scaling.
package units

// ByteSize counts bytes.
type ByteSize uint64

const (
	_ ByteSize = iota
	// CHANGE CODE BELOW THIS LINE
	KB ByteSize = 1 << iota
	// CHANGE CODE ABOVE THIS LINE
	MB
	GB
)
