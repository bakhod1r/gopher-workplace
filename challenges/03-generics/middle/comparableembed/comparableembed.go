// Package comparableembed — Gopher Workplace challenge.
package comparableembed

// Key is a comparable identifier restricted to ints and strings.
type Key interface {
	comparable
	~int | ~int64 | ~string
}

// Tally counts occurrences of each key.
// The Key constraint embeds comparable alongside a type set.
func Tally[K Key](s []K) map[K]int {
	// TODO(candidate): count each key.
	panic("not implemented")
}
