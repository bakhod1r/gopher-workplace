// Package multiembed — Gopher Workplace challenge.
package multiembed

// A provides a Name method.
type A struct{}

func (A) Name() string { return "A" }

// B provides a Name method.
type B struct{}

func (B) Name() string { return "B" }

// Collision embeds both A and B. Because both have a Name() method, calling
// c.Name() is a compile error (ambiguous). We must resolve it.
type Collision struct {
	A
	B
}

// Name resolves the collision by explicitly calling B's Name method.
//
// Examples:
//
//	Collision{}.Name() => "B"
func (c Collision) Name() string {
	// TODO(candidate): explicitly return the result of B's Name method.
	panic("not implemented")
}
