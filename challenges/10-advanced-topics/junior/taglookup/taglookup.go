// Package taglookup — Gopher Workplace challenge.
package taglookup

import "reflect"

// Tag returns the value of the given key in the named field's struct tag.
//
// The second result reports whether the field exists and carries that key.
//
// Examples:
//
//	Tag(row{}, "ID", "json") => "id", true
func Tag(v any, field, key string) (string, bool) {
	panic("not implemented")
}
