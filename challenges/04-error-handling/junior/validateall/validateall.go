// Package validateall — Gopher Workplace challenge.
package validateall

import "errors"

// ErrNegative reports a negative row value.
var ErrNegative = errors.New("negative value")

// Validate returns one error per negative entry in nums.
//
// Examples:
//
//	Validate([]int{1, -2, -3}) => [ErrNegative ErrNegative]
//	Validate([]int{1, 2})      => nil
func Validate(nums []int) []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
