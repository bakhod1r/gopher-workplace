// Package structsizeorder — Gopher Workplace challenge.
package structsizeorder

// Record must hold exactly these four fields:
//
//	ID      int64
//	Count   int32
//	Kind    int16
//	Enabled bool
//
// Declare them in the order that makes the struct as small as the alignment
// rules allow: 16 bytes on a 64-bit platform, not 24.
type Record struct {
	// declare the four fields here
}

// NewRecord builds a Record from its four values.
//
// Examples:
//
//	NewRecord(1, 2, 3, true) => Record{ID: 1, Count: 2, Kind: 3, Enabled: true}
func NewRecord(id int64, count int32, kind int16, enabled bool) Record {
	panic("not implemented")
}
