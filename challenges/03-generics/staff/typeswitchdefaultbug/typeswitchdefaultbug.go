// Package typeswitchdefaultbug — Gopher Workplace challenge.
package typeswitchdefaultbug

import (
	"strings"
)

// Normalize canonicalises v.
// Strings are trimmed and lower-cased; every other type is returned unchanged.
//
// Examples:
//
//	Normalize(" Hi ") => "hi"
//	Normalize(42) => 42
func Normalize[T any](v T) T {
	// CHANGE CODE BELOW THIS LINE
	switch x := any(v).(type) {
	case string:
		n := strings.ToLower(strings.TrimSpace(x))
		return any(n).(T)
	default:
		_ = x
		var zero T
		return zero
	}
	// CHANGE CODE ABOVE THIS LINE
}
