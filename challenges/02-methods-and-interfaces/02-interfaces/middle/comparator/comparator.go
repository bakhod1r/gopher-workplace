// Package comparator — Gopher Workplace challenge.
package comparator

// Record is one row.
type Record struct {
	Name string
	Age  int
}

// Comparator orders two records: <0, 0, or >0.
type Comparator interface {
	Compare(a, b Record) int
}

// ByName orders by name, ascending.
type ByName struct{}

// Compare orders by name.
func (ByName) Compare(a, b Record) int {
	// TODO(candidate): three-way string comparison.
	panic("not implemented")
}

// ByAge orders by age, ascending.
type ByAge struct{}

// Compare orders by age.
func (ByAge) Compare(a, b Record) int {
	// TODO(candidate): three-way int comparison.
	panic("not implemented")
}

// Reverse inverts any comparator.
type Reverse struct {
	Inner Comparator
}

// Compare returns the inverse of the wrapped comparison.
func (r Reverse) Compare(a, b Record) int {
	// TODO(candidate): negate the inner result.
	panic("not implemented")
}

// SortWith returns a sorted copy; the input is not modified.
//
// Examples:
//
//	SortWith(recs, ByAge{})                  => youngest first
//	SortWith(recs, Reverse{Inner: ByAge{}})  => oldest first
func SortWith(recs []Record, c Comparator) []Record {
	// TODO(candidate): copy, then sort the copy.
	panic("not implemented")
}
