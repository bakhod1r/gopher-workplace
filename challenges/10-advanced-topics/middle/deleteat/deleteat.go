// Package deleteat — Gopher Workplace challenge.
package deleteat

// Item is one stored element.
type Item struct {
	ID  int
	Pad [512]byte
}

// DeleteAt removes the element at index i, preserving the order of the
// rest, and returns the shortened slice.
//
// The vacated slot at the end must be cleared so the removed item stops
// being reachable through the backing array.
//
// Examples:
//
//	DeleteAt([]*Item{a, b, c}, 1) => a slice holding a and c
func DeleteAt(s []*Item, i int) []*Item {
	panic("not implemented")
}
