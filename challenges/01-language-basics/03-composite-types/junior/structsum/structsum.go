// Package structsum totals a field across a slice of structs.
package structsum

// Order is a line item.
type Order struct {
	Item  string
	Price int // cents
	Qty   int
}

// Total returns the sum of Price*Qty over all orders.
//
// TODO(candidate): range the slice and accumulate.
func Total(orders []Order) int {
	panic("not implemented")
}
