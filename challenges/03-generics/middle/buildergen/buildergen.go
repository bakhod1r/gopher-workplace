// Package buildergen — Gopher Workplace challenge.
package buildergen

// Builder collects values of T for fluent construction.
// Its zero value is ready to use.
type Builder[T any] struct {
	items []T
}

// With appends v and returns the builder for chaining.
func (b *Builder[T]) With(v T) *Builder[T] {
	// TODO(candidate): append v and return the receiver.
	panic("not implemented")
}

// WithAll appends every value and returns the builder.
func (b *Builder[T]) WithAll(vs ...T) *Builder[T] {
	// TODO(candidate): append all values and return the receiver.
	panic("not implemented")
}

// Build returns the collected values as a fresh slice.
func (b *Builder[T]) Build() []T {
	// TODO(candidate): return a copy of the collected values.
	panic("not implemented")
}
