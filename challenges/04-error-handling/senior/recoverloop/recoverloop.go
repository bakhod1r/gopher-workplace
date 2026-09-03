// Package recoverloop — Gopher Workplace challenge.
package recoverloop

import (
	"errors"
	"fmt"
)

// ErrPanic reports a recovered panic from one item.
var ErrPanic = errors.New("recovered panic")

// Process runs h for every item, isolating panics per item.
//
// Examples:
//
//	Process(nil, func(int) error { return nil }) => nil
func Process(items []int, h func(int) error) error {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, errors.Join
	panic("not implemented")
}
