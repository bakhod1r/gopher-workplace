// Package streamfilt — Gopher Workplace challenge.
package streamfilt

import "strings"

// Source yields lines until it is drained.
type Source interface {
	Next() (string, bool)
}

// Predicate decides whether a line is kept.
type Predicate interface {
	Match(line string) bool
}

// SliceSource streams a slice of lines.
type SliceSource struct {
	Lines []string
	pos   int
}

// Next returns the next line.
func (s *SliceSource) Next() (string, bool) {
	if s.pos >= len(s.Lines) {
		return "", false
	}
	line := s.Lines[s.pos]
	s.pos++
	return line, true
}

// Contains matches lines containing Sub.
type Contains struct {
	Sub string
}

// Match reports whether line contains Sub.
func (c Contains) Match(line string) bool {
	// TODO(candidate): substring test.
	panic("not implemented")
}

// MinLen matches lines of at least N bytes.
type MinLen struct {
	N int
}

// Match reports whether line is long enough.
func (m MinLen) Match(line string) bool {
	// TODO(candidate): length test.
	panic("not implemented")
}

// Not inverts a predicate.
type Not struct {
	Inner Predicate
}

// Match inverts the wrapped result.
func (n Not) Match(line string) bool {
	// TODO(candidate): negate.
	panic("not implemented")
}

// FilterStream drains src, keeping the lines that match p.
//
// Examples:
//
//	FilterStream(src, MinLen{N: 3}) => lines of 3 or more bytes
func FilterStream(src Source, p Predicate) []string {
	// TODO(candidate): stream, keep the matches.
	panic("not implemented")
}

var _ = strings.Contains
