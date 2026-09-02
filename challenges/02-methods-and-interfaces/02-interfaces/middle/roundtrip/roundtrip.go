// Package roundtrip — Gopher Workplace challenge.
package roundtrip

import "errors"

// ErrBadFormat reports unparsable input.
var ErrBadFormat = errors.New("bad format")

// Marshaler renders itself as text.
type Marshaler interface {
	Marshal() string
}

// Unmarshaler fills itself from text.
type Unmarshaler interface {
	Unmarshal(s string) error
}

// Codec does both.
type Codec interface {
	Marshaler
	Unmarshaler
}

// Record is an id/name pair.
type Record struct {
	ID   int
	Name string
}

// Marshal renders "<id>|<name>".
//
// Examples:
//
//	(&Record{ID: 1, Name: "a"}).Marshal() => "1|a"
func (r *Record) Marshal() string {
	// TODO(candidate): "<id>|<name>".
	panic("not implemented")
}

// Unmarshal parses "<id>|<name>".
//
// Examples:
//
//	r.Unmarshal("2|b")  => nil, Record{ID: 2, Name: "b"}
//	r.Unmarshal("oops") => ErrBadFormat
func (r *Record) Unmarshal(s string) error {
	// TODO(candidate): split, parse the id, reject bad input.
	panic("not implemented")
}

// RoundTrip marshals src, parses it into dst, and reports whether the
// re-marshalled text matches.
func RoundTrip(src, dst Codec) (bool, error) {
	// TODO(candidate): marshal, unmarshal, compare.
	panic("not implemented")
}
