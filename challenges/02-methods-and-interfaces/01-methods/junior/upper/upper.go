// Package upper — Gopher Workplace challenge.
package upper

import "strings"

// MyString is a string with methods.
type MyString string

// Upper returns the string in uppercase.
//
// Examples:
//
//	MyString("hello").Upper() => "HELLO"
//	MyString("Go").Upper()    => "GO"
func (s MyString) Upper() string {
	// TODO(candidate): implement this.
	_ = strings.ToUpper // hint
	panic("not implemented")
}
