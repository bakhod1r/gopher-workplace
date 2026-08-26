// Package sortby — Gopher Workplace challenge.
package sortby

import "sort"

// Person has a name and age.
type Person struct {
	Name string
	Age  int
}

// ByAge returns true if p is younger than other.
func (p Person) ByAge(other Person) bool {
	return p.Age < other.Age
}

// SortByField sorts people using the given comparison function.
// The comparison function has signature func(Person, Person) bool — it could
// be a method expression like Person.ByAge.
//
// Examples:
//
//	SortByField(people, Person.ByAge) // sorts by age ascending
func SortByField(people []Person, less func(Person, Person) bool) {
	// TODO(candidate): use sort.Slice with the less function.
	_ = sort.Slice // hint
	panic("not implemented")
}
