// Package grouprules — Gopher Workplace challenge.
package grouprules

import (
	"errors"
	"fmt"
)

// Rule is a named validation step.
type Rule struct {
	Name string
	Fn   func(string) error
}

// Check runs every rule against v, collecting named failures.
//
// Examples:
//
//	Check("abc", nil) => nil
func Check(v string, rules []Rule) error {
	// TODO(candidate): implement this.
	_, _ = fmt.Errorf, errors.Join
	panic("not implemented")
}
