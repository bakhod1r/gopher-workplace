// Package groupstructs groups orders by customer.
package groupstructs

// Order is a customer purchase.
type Order struct {
	Customer string
	Amount   int
}

// TotalByCustomer returns each customer's total order amount.
//
// TODO(candidate): accumulate into a map.
func TotalByCustomer(orders []Order) map[string]int {
	panic("not implemented")
}
