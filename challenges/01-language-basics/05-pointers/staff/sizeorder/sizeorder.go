// Package sizeorder defines a record whose field order must minimise padding to
// 16 bytes on a 64-bit platform. A planted bug orders fields bool,int64,bool,
// which pads to 24. Reorder to int64,bool,bool for 16.
package sizeorder

// Record must be exactly 16 bytes wide (64-bit).
// CHANGE CODE BELOW THIS LINE
type Record struct {
	A bool
	B int64
	C bool
}

// CHANGE CODE ABOVE THIS LINE
