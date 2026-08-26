// Package speaker — Gopher Workplace challenge.
package speaker

// Speaker is an interface for things that speak.
type Speaker interface {
	Speak() string
}

// Person speaks.
type Person struct{ Name string }

func (p Person) Speak() string {
	// TODO(candidate): return "Hi, I'm " + p.Name
	panic("not implemented")
}

// Robot speaks.
type Robot struct{ ID int }

func (r Robot) Speak() string {
	// TODO(candidate): return "I am robot" (ignore ID for simplicity)
	panic("not implemented")
}
