// Package formatter — Gopher Workplace challenge.
package formatter

import "fmt"

// Formatter is an interface for custom formatting.
type Formatter interface {
	Format() string
}

// Name holds first and last names.
type Name struct {
	First, Last string
}

// Format returns "Last, First".
func (n Name) Format() string {
	_ = fmt.Sprint
	// TODO(candidate): return n.Last + ", " + n.First
	panic("not implemented")
}

// FormatAll formats a slice of Formatters.
func FormatAll(fs []Formatter) []string {
	result := make([]string, len(fs))
	for i, f := range fs {
		result[i] = f.Format()
	}
	return result
}
