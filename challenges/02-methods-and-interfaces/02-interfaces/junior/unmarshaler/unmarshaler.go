// Package unmarshaler — Gopher Workplace challenge.
package unmarshaler

import "errors"

// ErrBadPair reports malformed input.
var ErrBadPair = errors.New("bad pair")

// Unmarshaler fills itself from wire text.
type Unmarshaler interface {
	Unmarshal(s string) error
}

// Pair is a key/value pair.
type Pair struct {
	Key   string
	Value string
}

// Unmarshal parses "<key>=<value>".
//
// Examples:
//
//	p := &Pair{}; p.Unmarshal("a=1")  => nil, Pair{Key: "a", Value: "1"}
//	p.Unmarshal("nope")               => ErrBadPair
func (p *Pair) Unmarshal(s string) error {
	// TODO(candidate): split on the first "=", reject input without one.
	panic("not implemented")
}

// UnmarshalAll parses every line, stopping at the first error.
func UnmarshalAll(lines []string) ([]Pair, error) {
	// TODO(candidate): parse each line into a Pair.
	panic("not implemented")
}
