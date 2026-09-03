// Package callfn — Gopher Workplace challenge.
package callfn

import (
	"errors"
	"reflect"
)

// ErrSignature is returned when fn does not match the expected shape.
var ErrSignature = errors.New("fn must take and return only ints")

// CallInts calls fn with args and returns its int results.
//
// fn must be a function taking exactly len(args) int parameters and
// returning only ints. Anything else is an error, not a panic.
//
// Examples:
//
//	CallInts(func(a, b int) int { return a + b }, 1, 2) => []int{3}
func CallInts(fn any, args ...int) ([]int, error) {
	panic("not implemented")
}
