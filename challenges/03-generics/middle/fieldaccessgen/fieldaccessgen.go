// Package fieldaccessgen — Gopher Workplace challenge.
package fieldaccessgen

// Book is a priced item.
type Book struct {
	Price int
}

// Coffee is a priced item.
type Coffee struct {
	Price int
}

// TotalPrice sums a projected field over any element type.
// The field is supplied as a function because a type parameter
// cannot have its fields accessed directly.
func TotalPrice[T any](items []T, price func(T) int) int {
	// TODO(candidate): sum the projected field.
	panic("not implemented")
}
