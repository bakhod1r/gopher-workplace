// Package encodeexact — Gopher Workplace challenge.
package encodeexact

import "strconv"

// Rec is one record to encode.
type Rec struct {
	ID   int
	Name string
}

// Encode renders each record as "id:name" separated by '\n'.
//
// The output's length is determined by the input, so it should be
// allocated once at exactly that size.
//
// Examples:
//
//	Encode([]Rec{{1, "a"}}) => []byte("1:a")
func Encode(recs []Rec) []byte {
	panic("not implemented")
}
