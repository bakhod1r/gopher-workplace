// Package eater — Gopher Workplace challenge.
package eater

// Eater reports whether it accepts a food.
type Eater interface {
	Eats(food string) bool
}

// Cow is a herbivore.
type Cow struct{}

// Eats reports whether the cow accepts food.
//
// Examples:
//
//	Cow{}.Eats("grass") => true
//	Cow{}.Eats("meat")  => false
func (c Cow) Eats(food string) bool {
	// TODO(candidate): grass only.
	panic("not implemented")
}

// Lion is a carnivore.
type Lion struct{}

// Eats reports whether the lion accepts food.
func (l Lion) Eats(food string) bool {
	// TODO(candidate): meat only.
	panic("not implemented")
}

// FeedableCount counts eaters that accept food.
func FeedableCount(es []Eater, food string) int {
	// TODO(candidate): count the accepting eaters.
	panic("not implemented")
}
