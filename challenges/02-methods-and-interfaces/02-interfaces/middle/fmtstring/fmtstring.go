// Package fmtstring — Gopher Workplace challenge.
package fmtstring

// Money is an amount in cents.
type Money int

// String renders the amount with two decimals.
//
// Examples:
//
//	Money(1234).String() => "12.34"
//	Money(5).String()    => "0.05"
//	Money(-250).String() => "-2.50"
func (m Money) String() string {
	// TODO(candidate): whole part, dot, two-digit cents.
	panic("not implemented")
}

// Level is a log severity.
type Level int

// Known levels.
const (
	Debug Level = iota
	Info
	Error
)

// String names the level, or "LEVEL(n)" when unknown.
func (l Level) String() string {
	// TODO(candidate): map the level to a name.
	panic("not implemented")
}

// Line renders "[<level>] <msg>: <money>".
func Line(l Level, msg string, m Money) string {
	// TODO(candidate): use the Stringer implementations.
	panic("not implemented")
}
