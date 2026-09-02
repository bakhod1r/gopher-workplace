// Package stringerifc — Gopher Workplace challenge.
package stringerifc

import (
	"fmt"
)

// Tag is a label that knows how to describe itself.
type Tag string

// String renders the tag.
func (t Tag) String() string { return "tag:" + string(t) }

// Describe returns the String() of every element.
// It uses a type parameter, so the caller keeps a typed slice.
func Describe[T fmt.Stringer](items []T) []string {
	// TODO(candidate): call String on each element.
	panic("not implemented")
}

// DescribeAny does the same for a slice of interface values.
// It is provided so you can compare the two shapes.
func DescribeAny(items []fmt.Stringer) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		out = append(out, it.String())
	}
	return out
}
