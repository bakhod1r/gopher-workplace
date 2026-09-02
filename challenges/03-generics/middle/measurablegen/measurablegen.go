// Package measurablegen — Gopher Workplace challenge.
package measurablegen

// Measurable is anything that can report itself as a number.
type Measurable interface {
	Value() float64
}

// Reading is a sensor sample.
type Reading struct {
	V float64
}

// Value reports the sample's magnitude.
func (r Reading) Value() float64 { return r.V }

// Heaviest returns the sample with the largest value and true.
func Heaviest[T Measurable](s []T) (T, bool) {
	// TODO(candidate): track the sample with the largest Value().
	panic("not implemented")
}
