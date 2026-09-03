// Package embeddedfields — Gopher Workplace challenge.
package embeddedfields

import "reflect"

// Paths returns the dotted path of every exported leaf field of v,
// descending through embedded structs.
//
// An embedded struct contributes its fields' paths under its own name;
// named struct fields are not descended into.
//
// Examples:
//
//	Paths(User{}) => []string{"Base.ID", "Name"}
func Paths(v any) []string {
	panic("not implemented")
}
