// Package buildcapbug — Gopher Workplace challenge.
package buildcapbug

// Builder accumulates elements of any type.
type Builder[T any] struct {
	items []T
}

// Build returns the accumulated elements.
// The returned slice is independent of the builder.
func (b *Builder[T]) Build() []T {
	// CHANGE CODE BELOW THIS LINE
	return b.items
	// CHANGE CODE ABOVE THIS LINE
}

// Add appends v and returns the builder for chaining.
func (b *Builder[T]) Add(v T) *Builder[T] {
	b.items = append(b.items, v)
	return b
}
