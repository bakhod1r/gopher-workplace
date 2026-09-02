// Package sumpositive — Gopher Workplace challenge.
package sumpositive

import "errors"

// ErrNegativeValue reports a negative entry.
var ErrNegativeValue = errors.New("negative value")

// SumPositive totals nums, rejecting any negative entry.
//
// Examples:
//
//	SumPositive([]int{1, 2, 3}) => 6, nil
//	SumPositive([]int{1, -2})   => 0, ErrNegativeValue
func SumPositive(nums []int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
