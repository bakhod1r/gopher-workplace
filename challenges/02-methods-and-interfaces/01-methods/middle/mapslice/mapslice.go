// Package mapslice — Gopher Workplace challenge.
package mapslice

import "fmt"

// IntList is a slice of ints.
type IntList []int

// ToString returns a new slice of strings where each int is formatted as string.
//
// Examples:
//
//	IntList{1, 2, 3}.ToString() => {"1", "2", "3"}
func (l IntList) ToString() []string {
	// TODO(candidate): map ints to strings.
	_ = fmt.Sprint // hint
	panic("not implemented")
}
