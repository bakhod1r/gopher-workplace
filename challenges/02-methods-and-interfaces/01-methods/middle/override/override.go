// Package override — Gopher Workplace challenge.
package override

import "fmt"

// Animal has a species name.
type Animal struct {
	Species string
}

// String returns "Animal(<Species>)".
func (a Animal) String() string {
	return fmt.Sprintf("Animal(%s)", a.Species)
}

// Pet embeds Animal and adds a nickname.
type Pet struct {
	Animal
	Nickname string
}

// String overrides the promoted Animal.String() to include the nickname.
// Returns "Pet(<Nickname>, <Species>)".
//
// Examples:
//
//	Pet{Animal{"Cat"}, "Whiskers"}.String() => "Pet(Whiskers, Cat)"
func (p Pet) String() string {
	// TODO(candidate): override the promoted String with a version that
	// includes Nickname. You can still access the embedded method via
	// p.Animal.String() if you want, but it's not required.
	panic("not implemented")
}
