// Package rangecopymutate applies a discount to every order. A planted bug
// mutates the range copy, so nothing changes.
package rangecopymutate

// Order has a price in cents.
type Order struct {
	Price int
}

// Discount reduces every order's price by pct percent (integer), in place.
func Discount(orders []Order, pct int) {
	// CHANGE CODE BELOW THIS LINE
	for _, o := range orders {
		o.Price -= o.Price * pct / 100
	}
	// CHANGE CODE ABOVE THIS LINE
}
