// Package encoder — Gopher Workplace challenge.
package encoder

import "strings"

// CSVEncoder encodes records as CSV rows.
type CSVEncoder struct {
	Rows []string
}

// Encode adds a row from the given fields, separated by commas.
//
// Examples:
//
//	e := &CSVEncoder{}
//	e.Encode("a", "b", "c") // e.Rows == ["a,b,c"]
//	e.Encode("1", "2")      // e.Rows == ["a,b,c", "1,2"]
func (e *CSVEncoder) Encode(fields ...string) {
	// TODO(candidate): join fields with commas, append to Rows.
	_ = strings.Join // hint
	panic("not implemented")
}
