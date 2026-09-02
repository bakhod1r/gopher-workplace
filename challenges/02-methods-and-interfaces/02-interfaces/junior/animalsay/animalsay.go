// Package animalsay — Gopher Workplace challenge.
package animalsay

// Animal makes a sound.
type Animal interface {
	Sound() string
}

// Dog barks.
type Dog struct{}

// Sound returns the dog's noise.
//
// Examples:
//
//	Dog{}.Sound() => "Woof!"
func (d Dog) Sound() string {
	// TODO(candidate): "Woof!".
	panic("not implemented")
}

// Cat meows.
type Cat struct{}

// Sound returns the cat's noise.
func (c Cat) Sound() string {
	// TODO(candidate): "Meow!".
	panic("not implemented")
}

// MakeNoise returns the sound of any animal.
func MakeNoise(a Animal) string {
	// TODO(candidate): ask the animal.
	panic("not implemented")
}

// Chorus joins every sound with a single space.
func Chorus(as []Animal) string {
	// TODO(candidate): space-separated sounds.
	panic("not implemented")
}
