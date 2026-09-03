// Package stringinterner — Gopher Workplace challenge.
package stringinterner

// Interner maps equal strings onto one shared instance, so a million repeated
// labels cost one string's worth of memory instead of a million. The zero
// value is ready to use.
type Interner struct {
	seen   map[string]string
	hits   int
	misses int
}

// Intern returns the canonical instance of s: the first one interned wins,
// and every later equal string resolves to it.
//
// Examples:
//
//	in.Intern("a") => "a"
func (in *Interner) Intern(s string) string {
	panic("not implemented")
}

// InternBytes interns the string form of b without keeping a reference to b,
// which the caller may reuse. It must not allocate when the string is already
// known — the compiler elides the conversion in a map lookup.
//
// Examples:
//
//	in.InternBytes([]byte("a")) => "a"
func (in *Interner) InternBytes(b []byte) string {
	panic("not implemented")
}

// Stats reports how many lookups hit an existing entry and how many added one.
//
// Examples:
//
//	in.Stats() => 3, 1
func (in *Interner) Stats() (hits, misses int) {
	panic("not implemented")
}

// Len reports how many distinct strings are held.
//
// Examples:
//
//	in.Len() => 1
func (in *Interner) Len() int {
	panic("not implemented")
}
