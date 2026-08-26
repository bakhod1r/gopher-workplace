// Package jsonunmarshal — Gopher Workplace challenge.
package jsonunmarshal

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Money represents an amount in cents.
type Money struct {
	Cents int
}

// UnmarshalJSON parses a JSON string like "$10.50" into cents.
// Implements json.Unmarshaler.
//
// Examples:
//
//	`"$10.50"` => Money{1050}
func (m *Money) UnmarshalJSON(data []byte) error {
	// TODO(candidate): parse data (with quotes), strip quotes and $,
	// convert dollars and cents to an int, store in m.Cents.
	_ = json.Unmarshal // hint
	_ = fmt.Sprint
	_ = strings.HasPrefix
	panic("not implemented")
}
