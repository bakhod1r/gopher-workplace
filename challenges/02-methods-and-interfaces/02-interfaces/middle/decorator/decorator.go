// Package decorator — Gopher Workplace challenge.
package decorator

// Pricer computes a price in minor units.
type Pricer interface {
	Price() int
}

// Base is the raw price.
type Base struct {
	Amount int
}

// Price returns the raw amount.
func (b Base) Price() int {
	// TODO(candidate): return Amount.
	panic("not implemented")
}

// Discount takes a percentage off the wrapped price.
type Discount struct {
	Inner   Pricer
	Percent int
}

// Price applies the discount, truncating.
//
// Examples:
//
//	Discount{Inner: Base{Amount: 100}, Percent: 10}.Price() => 90
func (d Discount) Price() int {
	// TODO(candidate): subtract Percent% of the inner price.
	panic("not implemented")
}

// Tax adds a percentage to the wrapped price.
type Tax struct {
	Inner   Pricer
	Percent int
}

// Price applies the tax, truncating.
func (t Tax) Price() int {
	// TODO(candidate): add Percent% of the inner price.
	panic("not implemented")
}

// Layer wraps a pricer in one more decoration.
type Layer func(Pricer) Pricer

// Wrap applies each layer in order: the first listed wraps the base first.
func Wrap(p Pricer, layers ...Layer) Pricer {
	// TODO(candidate): fold the layers over p.
	panic("not implemented")
}
