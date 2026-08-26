// Package flyer — Gopher Workplace challenge.
package flyer

// Flyer is an interface for things that fly.
type Flyer interface {
	Fly() string
}

// Bird flies.
type Bird struct{ Species string }

func (b Bird) Fly() string {
	// TODO(candidate): return b.Species + " flies"
	panic("not implemented")
}

// Plane flies.
type Plane struct{}

func (p Plane) Fly() string {
	// TODO(candidate): return "plane flies"
	panic("not implemented")
}
