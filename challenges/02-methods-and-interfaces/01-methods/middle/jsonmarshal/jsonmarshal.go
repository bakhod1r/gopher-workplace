// Package jsonmarshal — Gopher Workplace challenge.
package jsonmarshal

import "fmt"

// Money represents an amount in cents.
type Money struct {
	Cents int
}

// MarshalJSON returns the JSON representation as a quoted dollar string.
// Implements json.Marshaler.
//
// Examples:
//
//	Money{1050}.MarshalJSON() => []byte(`"$10.50"`), nil
//	Money{99}.MarshalJSON()   => []byte(`"$0.99"`), nil
//	Money{0}.MarshalJSON()    => []byte(`"$0.00"`), nil
func (m Money) MarshalJSON() ([]byte, error) {
	// TODO(candidate): format as "$X.YY" and return as JSON string bytes.
	_ = fmt.Sprintf // hint
	panic("not implemented")
}
