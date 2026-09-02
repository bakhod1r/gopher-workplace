// Package combineconstraint — Gopher Workplace challenge.
package combineconstraint

// Number is the numeric set.
type Number interface {
	~int | ~int64 | ~float64
}

// Text is the string-like set.
type Text interface {
	~string
}

// NumOrText is the union of the two sets above.
type NumOrText interface {
	Number | Text
}

// Render formats every element with fmt.Sprint.
// Its constraint is the union of two other constraints.
func Render[T NumOrText](s []T) []string {
	// TODO(candidate): format each element.
	panic("not implemented")
}
