// Package structpadding lays out a record compactly. A planted bug orders fields
// so padding bloats the struct.
package structpadding

// Record should be laid out to minimize size via field ordering. On a 64-bit
// platform the minimal size is 16 bytes.
type Record struct {
	// CHANGE CODE BELOW THIS LINE
	A bool
	B int64
	C bool
	// CHANGE CODE ABOVE THIS LINE
}
