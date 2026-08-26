// Package describe — Gopher Workplace challenge.
package describe

import "fmt"

// Circle has a radius.
type Circle struct {
	Radius float64
}

// Describe returns "Circle(radius=R)" where R is the radius.
//
// Examples:
//
//	Circle{5}.Describe()   => "Circle(radius=5)"
//	Circle{3.5}.Describe() => "Circle(radius=3.5)"
func (c Circle) Describe() string {
	// TODO(candidate): implement this.
	_ = fmt.Sprintf // hint
	panic("not implemented")
}

// Square has a side length.
type Square struct {
	Side float64
}

// Describe returns "Square(side=S)" where S is the side.
//
// Examples:
//
//	Square{4}.Describe()   => "Square(side=4)"
func (s Square) Describe() string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
