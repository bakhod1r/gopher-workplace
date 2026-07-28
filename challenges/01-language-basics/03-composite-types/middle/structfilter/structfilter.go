// Package structfilter filters records by a field.
package structfilter

// User has a name and an active flag.
type User struct {
	Name   string
	Active bool
}

// ActiveNames returns the names of active users, in order.
//
// TODO(candidate): filter and project.
func ActiveNames(users []User) []string {
	panic("not implemented")
}
