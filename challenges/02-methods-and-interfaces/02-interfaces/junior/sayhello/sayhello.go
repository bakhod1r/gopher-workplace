// Package sayhello — Gopher Workplace challenge.
package sayhello

// Greeter is an interface for things that can greet.
type Greeter interface {
	Greet() string
}

// English greets in English.
type English struct{}

// Greet returns an English greeting.
//
// Examples:
//
//	English{}.Greet() => "Hello!"
func (e English) Greet() string {
	// TODO(candidate): return "Hello!"
	panic("not implemented")
}

// Uzbek greets in Uzbek.
type Uzbek struct{}

// Greet returns an Uzbek greeting.
//
// Examples:
//
//	Uzbek{}.Greet() => "Salom!"
func (u Uzbek) Greet() string {
	// TODO(candidate): return "Salom!"
	panic("not implemented")
}

// SayHello takes any Greeter and returns its greeting.
func SayHello(g Greeter) string {
	return g.Greet()
}
