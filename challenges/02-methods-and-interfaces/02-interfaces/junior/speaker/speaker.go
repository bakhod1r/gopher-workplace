// Package speaker — Gopher Workplace challenge.
package speaker

// Speaker says one line about itself.
type Speaker interface {
	Speak() string
}

// Person introduces itself by name.
type Person struct {
	Name string
}

// Speak returns the person's line.
//
// Examples:
//
//	Person{Name: "Go"}.Speak() => "Hi, I'm Go"
func (p Person) Speak() string {
	// TODO(candidate): "Hi, I'm <Name>".
	panic("not implemented")
}

// Robot has a fixed line.
type Robot struct {
	ID int
}

// Speak returns the robot's line.
func (r Robot) Speak() string {
	// TODO(candidate): "I am robot".
	panic("not implemented")
}

// Introduce returns what s says.
func Introduce(s Speaker) string {
	// TODO(candidate): ask the speaker.
	panic("not implemented")
}
