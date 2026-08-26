// Package decoder — Gopher Workplace challenge.
package decoder

import "strings"

// CSVDecoder reads CSV rows.
type CSVDecoder struct {
	rows []string
	pos  int
}

// NewCSVDecoder creates a decoder from raw CSV lines.
func NewCSVDecoder(lines []string) *CSVDecoder {
	return &CSVDecoder{rows: lines}
}

// Next advances to the next row. Returns false when exhausted.
func (d *CSVDecoder) Next() bool {
	// TODO(candidate): advance and check bounds.
	panic("not implemented")
}

// Fields returns the current row split by commas.
//
// Examples:
//
//	d := NewCSVDecoder([]string{"a,b,c", "1,2"})
//	d.Next()   => true
//	d.Fields() => ["a", "b", "c"]
//	d.Next()   => true
//	d.Fields() => ["1", "2"]
func (d *CSVDecoder) Fields() []string {
	// TODO(candidate): split the current row by commas.
	_ = strings.Split // hint
	panic("not implemented")
}
