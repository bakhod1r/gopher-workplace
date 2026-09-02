// Package constraintifcmethod — Gopher Workplace challenge.
package constraintifcmethod

// Priced is anything that knows its price in cents.
type Priced interface {
	Cents() int
}

// Book is a priced item.
type Book struct {
	Price int
}

// Cents returns the book's price.
func (b Book) Cents() int { return b.Price }

// Coffee is a priced item.
type Coffee struct {
	Price int
}

// Cents returns the coffee's price.
func (c Coffee) Cents() int { return c.Price }

// TotalCents sums the cents of every item.
// The constraint requires both a type set and a method.
func TotalCents[T Priced](items []T) int {
	// TODO(candidate): sum the Cents of each item.
	panic("not implemented")
}
