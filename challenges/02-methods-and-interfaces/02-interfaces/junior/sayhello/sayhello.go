// Package sayhello — Gopher Workplace challenge.
package sayhello

// Greeter returns a greeting word.
type Greeter interface {
	Hello() string
}

// English greets in English.
type English struct{}

// Hello returns the English greeting.
func (e English) Hello() string {
	// TODO(candidate): "Hello".
	panic("not implemented")
}

// Uzbek greets in Uzbek.
type Uzbek struct{}

// Hello returns the Uzbek greeting.
func (u Uzbek) Hello() string {
	// TODO(candidate): "Salom".
	panic("not implemented")
}

// Greet builds "<greeting>, <name>".
//
// Examples:
//
//	Greet(Uzbek{}, "Ali") => "Salom, Ali"
func Greet(g Greeter, name string) string {
	// TODO(candidate): combine greeting and name.
	panic("not implemented")
}
