// Package status enumerates order states. A planted zero-value trap misclassifies.
package status

// Status is an order lifecycle state.
type Status int

const (
	// CHANGE CODE BELOW THIS LINE
	Pending Status = iota
	// CHANGE CODE ABOVE THIS LINE
	Shipped
	Delivered
)

// IsKnown reports whether s is one of the defined states.
// A brand-new zero-valued Status must be reported as unknown.
func IsKnown(s Status) bool {
	return s == Pending || s == Shipped || s == Delivered
}
