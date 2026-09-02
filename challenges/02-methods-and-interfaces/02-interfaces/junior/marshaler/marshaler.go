// Package marshaler — Gopher Workplace challenge.
package marshaler

// Marshaler renders itself in wire format.
type Marshaler interface {
	Marshal() string
}

// Point is a 2D coordinate.
type Point struct {
	X, Y int
}

// Marshal renders the point.
//
// Examples:
//
//	Point{X: 1, Y: 2}.Marshal() => "1,2"
func (p Point) Marshal() string {
	// TODO(candidate): "<X>,<Y>".
	panic("not implemented")
}

// Label is a plain text tag.
type Label string

// Marshal renders the label.
func (l Label) Marshal() string {
	// TODO(candidate): the label itself.
	panic("not implemented")
}

// MarshalAll renders every value, in order.
func MarshalAll(ms []Marshaler) []string {
	// TODO(candidate): collect Marshal() results.
	panic("not implemented")
}
