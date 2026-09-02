// Package longesttext — Gopher Workplace challenge.
package longesttext

// Text is the set of string-like types.
type Text interface {
	~string
}

// Label is a short display string.
type Label string

// Longest returns the longest element of s and true.
// On a tie the earlier element wins.
// It returns the zero value and false for an empty slice.
//
// Examples:
//
//	Longest([]string{"a", "bbb"}) => "bbb", true
func Longest[T Text](s []T) (T, bool) {
	// TODO(candidate): track the element with the greatest length.
	panic("not implemented")
}
