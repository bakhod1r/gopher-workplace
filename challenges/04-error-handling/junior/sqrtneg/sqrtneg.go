// Package sqrtneg — Gopher Workplace challenge.
package sqrtneg

import (
	"errors"
	"math"
)

// ErrNegative reports a negative input.
var ErrNegative = errors.New("negative input")

// Sqrt returns the square root of x, or ErrNegative when x < 0.
//
// Examples:
//
//	Sqrt(9)  => 3, nil
//	Sqrt(-1) => 0, ErrNegative
func Sqrt(x float64) (float64, error) {
	// TODO(candidate): implement this.
	_ = math.Sqrt
	panic("not implemented")
}
