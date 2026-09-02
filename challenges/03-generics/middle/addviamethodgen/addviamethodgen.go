// Package addviamethodgen — Gopher Workplace challenge.
package addviamethodgen

// Adder is a type that can add another T to itself.
type Adder[T any] interface {
	Plus(T) T
}

// Money is an amount in minor units.
type Money struct {
	Cents int
}

// Plus adds two amounts.
func (m Money) Plus(o Money) Money { return Money{Cents: m.Cents + o.Cents} }

// SumAll adds values using their Plus method, since Go has no
// operator overloading for user types.
func SumAll[T Adder[T]](s []T) T {
	// TODO(candidate): fold the values with their Plus method.
	panic("not implemented")
}
