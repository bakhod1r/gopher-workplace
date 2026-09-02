// Package groupby — Gopher Workplace challenge.
package groupby

import "strconv"

// Keyer extracts the grouping key of a value.
type Keyer interface {
	Key(s string) string
}

// ByFirstLetter groups by the first byte; the empty string groups under "".
type ByFirstLetter struct{}

// Key returns the first letter.
func (ByFirstLetter) Key(s string) string {
	// TODO(candidate): first byte, or "" for an empty string.
	panic("not implemented")
}

// ByLength groups by decimal length.
type ByLength struct{}

// Key returns the length as text.
func (ByLength) Key(s string) string {
	// TODO(candidate): decimal length.
	panic("not implemented")
}

// Group buckets values by key in one pass, keeping input order per bucket.
//
// Examples:
//
//	Group([]string{"apple", "avocado", "beet"}, ByFirstLetter{})
//	  => {"a": ["apple", "avocado"], "b": ["beet"]}
func Group(values []string, k Keyer) map[string][]string {
	// TODO(candidate): one pass, append into the bucket.
	panic("not implemented")
}

// SortedKeys returns the grouping's keys in sorted order.
func SortedKeys(groups map[string][]string) []string {
	// TODO(candidate): collect and sort.
	panic("not implemented")
}

var _ = strconv.Itoa
