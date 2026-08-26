// Package isadult — Gopher Workplace challenge.
package isadult

// Person holds a name and age.
type Person struct {
	Name string
	Age  int
}

// IsAdult returns true if the person's age is 18 or older.
//
// Examples:
//
//	Person{"Alice", 25}.IsAdult() => true
//	Person{"Bob", 17}.IsAdult()   => false
func (p Person) IsAdult() bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
