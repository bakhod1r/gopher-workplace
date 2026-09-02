// Package parseintorgen — Gopher Workplace challenge.
package parseintorgen

// Integer is the set of signed integer types used here.
type Integer interface {
	~int | ~int64
}

// Retries counts attempts.
type Retries int

// ParseOr parses s as an integer, returning def when it cannot.
func ParseOr[T Integer](s string, def T) T {
	// TODO(candidate): parse, falling back to def on failure.
	panic("not implemented")
}
