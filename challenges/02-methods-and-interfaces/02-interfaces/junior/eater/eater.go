// Package eater — Gopher Workplace challenge.
package eater

// Eater is an interface for things that eat.
type Eater interface {
	Eat(food string) string
}

// Human eats.
type Human struct{ Name string }

func (h Human) Eat(food string) string {
	// TODO(candidate): return h.Name + " eats " + food
	panic("not implemented")
}
