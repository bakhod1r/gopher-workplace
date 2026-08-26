// Package animalsay — Gopher Workplace challenge.
package animalsay

// Animal is an interface for things that make sounds.
type Animal interface {
	Sound() string
}

// Dog says woof.
type Dog struct{}

func (d Dog) Sound() string {
	// TODO(candidate): return "Woof!"
	panic("not implemented")
}

// Cat says meow.
type Cat struct{}

func (c Cat) Sound() string {
	// TODO(candidate): return "Meow!"
	panic("not implemented")
}

// MakeNoise returns the sound of any animal.
func MakeNoise(a Animal) string {
	return a.Sound()
}
