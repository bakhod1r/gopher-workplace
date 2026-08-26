// Package fullname — Gopher Workplace challenge.
package fullname

// Person holds a first and last name.
type Person struct {
	First string
	Last  string
}

// FullName returns "First Last".
//
// Examples:
//
//	Person{"Go", "Gopher"}.FullName() => "Go Gopher"
//	Person{"", "Doe"}.FullName()      => " Doe"
func (p Person) FullName() string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
