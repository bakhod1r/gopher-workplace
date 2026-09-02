// Package schemaloader - Gopher Workplace challenge.
package schemaloader

import "sync"

// Loader parses a GraphQL schema exactly once and caches the outcome.
type Loader struct {
	once     sync.Once
	parse    func() (string, error)
	schema   string
	err      error
	attempts int
}

// NewLoader returns a Loader that calls parse on first use.
func NewLoader(parse func() (string, error)) *Loader {
	return &Loader{parse: parse}
}

// Load returns the parsed schema, parsing on the first call only.
//
// Examples:
//
//	l := NewLoader(func() (string, error) { return "schema", nil }); l.Load() => "schema", nil
//	l.Load(); l.Load()                                                        => same pair, parsed once
func (l *Loader) Load() (string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Attempts reports how many times the parse function ran.
//
// Examples:
//
//	l.Load(); l.Load(); l.Attempts() => 1
func (l *Loader) Attempts() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
