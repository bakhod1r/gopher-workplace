// Package swimmer — Gopher Workplace challenge.
package swimmer

// Swimmer describes its swimming.
type Swimmer interface {
	Swim() string
}

// Fish swims by name.
type Fish struct {
	Name string
}

// Swim returns the fish's line.
//
// Examples:
//
//	Fish{Name: "nemo"}.Swim() => "nemo swims"
func (f Fish) Swim() string {
	// TODO(candidate): "<Name> swims".
	panic("not implemented")
}

// Duck also swims.
type Duck struct{}

// Swim returns the duck's line.
func (d Duck) Swim() string {
	// TODO(candidate): "duck swims".
	panic("not implemented")
}

// SwimAll returns one line per swimmer, in order.
func SwimAll(ss []Swimmer) []string {
	// TODO(candidate): collect Swim() results.
	panic("not implemented")
}
