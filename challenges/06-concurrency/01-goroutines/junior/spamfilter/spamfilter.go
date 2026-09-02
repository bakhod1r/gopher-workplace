// Package spamfilter — Gopher Workplace challenge.
package spamfilter

// Flagged marks every message that contains the banned phrase.
//
// Examples:
//
//	Flagged([]string{"buy now", "hello"}, "buy")  => [true false]
//	Flagged([]string{"hello"}, "buy")             => [false]
//	Flagged(nil, "buy")                           => []
func Flagged(messages []string, banned string) []bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
