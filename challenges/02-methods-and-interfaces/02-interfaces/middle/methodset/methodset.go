// Package methodset — Gopher Workplace challenge.
package methodset

// Named reports a name.
type Named interface {
	Name() string
}

// Renamer can change its name.
type Renamer interface {
	Rename(s string)
}

// Value uses a value receiver.
type Value struct {
	N string
}

// Name returns the name.
func (v Value) Name() string {
	// TODO(candidate): return N.
	panic("not implemented")
}

// Pointer uses pointer receivers.
type Pointer struct {
	N string
}

// Name returns the name.
func (p *Pointer) Name() string {
	// TODO(candidate): return N.
	panic("not implemented")
}

// Rename replaces the name.
func (p *Pointer) Rename(s string) {
	// TODO(candidate): set N.
	panic("not implemented")
}

// Names collects the name of every element.
func Names(ns []Named) []string {
	// TODO(candidate): one name per element, in order.
	panic("not implemented")
}

// Satisfies reports whether v satisfies Renamer.
//
// Examples:
//
//	Satisfies(&Pointer{}) => true
//	Satisfies(Pointer{})  => false
func Satisfies(v any) bool {
	// TODO(candidate): interface assertion.
	panic("not implemented")
}
