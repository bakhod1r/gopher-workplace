// Package comparablepitfall — Gopher Workplace challenge.
package comparablepitfall

// Point is a comparable struct: every field is comparable.
type Point struct {
	X int
	Y int
}

// CountDistinct returns how many distinct values s holds.
func CountDistinct[T comparable](s []T) int {
	// TODO(candidate): count the distinct values.
	panic("not implemented")
}
