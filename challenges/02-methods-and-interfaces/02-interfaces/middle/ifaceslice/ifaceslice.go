// Package ifaceslice — Gopher Workplace challenge.
package ifaceslice

// Entity has an id.
type Entity interface {
	ID() string
}

// User is an account record.
type User struct {
	ID_ string
}

// ID returns the user id.
func (u User) ID() string {
	// TODO(candidate): return ID_.
	panic("not implemented")
}

// Order is a purchase record.
type Order struct {
	ID_ string
}

// ID returns the order id.
func (o Order) ID() string {
	// TODO(candidate): return ID_.
	panic("not implemented")
}

// ToEntities converts users into entities, preserving order.
//
// Examples:
//
//	ToEntities([]User{{ID_: "u1"}}) => []Entity{User{ID_: "u1"}}
func ToEntities(users []User) []Entity {
	// TODO(candidate): element-by-element conversion, preallocated.
	panic("not implemented")
}

// IDs returns the id of every entity, in order.
func IDs(es []Entity) []string {
	// TODO(candidate): collect ids.
	panic("not implemented")
}
