// Package treestring — Gopher Workplace challenge.
package treestring

import (
	"errors"
	"strings"
)

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Tree renders err as an indented tree, one line per error.
//
// Examples:
//
//	Tree(nil)  => ""
//	Tree(ErrA) => "a"
func Tree(err error) string {
	// TODO(candidate): implement this.
	_ = strings.Repeat
	panic("not implemented")
}
