// Package structclone — Gopher Workplace challenge.
package structclone

import "reflect"

// Clone returns a copy of the struct v, as a value of the same type.
//
// The copy is shallow: fields are assigned, so slices and maps inside it
// still share their storage with v.
//
// Examples:
//
//	Clone(pt{1, 2}) => pt{1, 2}, a distinct value
func Clone(v any) any {
	panic("not implemented")
}
