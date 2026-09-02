// Package stringer — Gopher Workplace challenge.
package stringer

// Stringer renders itself as text (same shape as fmt.Stringer).
type Stringer interface {
	String() string
}

// Color is a colour code.
type Color int

// Known colours.
const (
	Red Color = iota
	Green
	Blue
)

// String names the colour, or "unknown".
//
// Examples:
//
//	Red.String()      => "red"
//	Color(9).String() => "unknown"
func (c Color) String() string {
	// TODO(candidate): map the code to a name.
	panic("not implemented")
}

// Temp is a temperature in Celsius.
type Temp int

// String renders "<n>C".
func (t Temp) String() string {
	// TODO(candidate): number followed by "C".
	panic("not implemented")
}

// Print returns the text form of s.
func Print(s Stringer) string {
	// TODO(candidate): ask the value.
	panic("not implemented")
}
