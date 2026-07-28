// Package sortstructs sorts records by a field.
package sortstructs

// Person has a name and age.
type Person struct {
	Name string
	Age  int
}

// ByAge returns a new slice sorted by Age ascending; ties keep Name ascending.
// The input is not modified.
//
// TODO(candidate): copy, then sort.Slice with a compound comparator.
func ByAge(people []Person) []Person {
	panic("not implemented")
}
