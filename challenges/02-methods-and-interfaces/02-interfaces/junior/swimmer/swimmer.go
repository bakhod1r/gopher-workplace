// Package swimmer — Gopher Workplace challenge.
package swimmer

// Swimmer is an interface for things that swim.
type Swimmer interface {
	Swim() string
}

// Fish swims.
type Fish struct{ Name string }

func (f Fish) Swim() string {
	// TODO(candidate): return f.Name + " swims"
	panic("not implemented")
}

// Duck also swims.
type Duck struct{}

func (d Duck) Swim() string {
	// TODO(candidate): return "duck swims"
	panic("not implemented")
}
