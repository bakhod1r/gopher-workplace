// Package flyer — Gopher Workplace challenge.
package flyer

// Flyer reports its cruising altitude in metres.
type Flyer interface {
	Altitude() int
}

// Bird cruises in metres.
type Bird struct {
	Meters int
}

// Altitude returns metres.
func (b Bird) Altitude() int {
	// TODO(candidate): already metres.
	panic("not implemented")
}

// Jet cruises in feet.
type Jet struct {
	Feet int
}

// Altitude converts feet to metres, truncating.
//
// Examples:
//
//	Jet{Feet: 30000}.Altitude() => 9144
func (j Jet) Altitude() int {
	// TODO(candidate): feet -> metres.
	panic("not implemented")
}

// Highest returns the greatest altitude, or 0 for an empty slice.
func Highest(fs []Flyer) int {
	// TODO(candidate): track the running maximum.
	panic("not implemented")
}
