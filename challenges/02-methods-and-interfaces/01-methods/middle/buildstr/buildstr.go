// Package buildstr — Gopher Workplace challenge.
package buildstr

import "strings"

// Builder accumulates parts and joins them.
type Builder struct {
	parts []string
	sep   string
}

// NewBuilder creates a Builder with the given separator.
func NewBuilder(sep string) *Builder {
	return &Builder{sep: sep}
}

// Add appends a part and returns the builder for chaining.
//
// Examples:
//
//	NewBuilder(", ").Add("a").Add("b").Add("c").Build() => "a, b, c"
func (b *Builder) Add(part string) *Builder {
	// TODO(candidate): append part and return b.
	panic("not implemented")
}

// Build joins all parts with the separator.
func (b *Builder) Build() string {
	return strings.Join(b.parts, b.sep)
}
