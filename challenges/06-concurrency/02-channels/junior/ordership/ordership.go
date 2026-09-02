// Package ordership — Gopher Workplace challenge.
package ordership

// ShipOrderIDs sends every order id on out, in order, then closes out.
//
// Examples:
//
//	ShipOrderIDs(out, []int{101, 102}) => out yields 101, 102 then closed
//	ShipOrderIDs(out, nil)             => out closed immediately
//	ShipOrderIDs(out, []int{7})        => out yields 7 then closed
func ShipOrderIDs(out chan<- int, ids []int) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
