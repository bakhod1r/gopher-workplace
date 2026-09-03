// Package tagvalidate — Gopher Workplace challenge.
package tagvalidate

import "reflect"

// Validate returns the problems with v's `col` tags, in field order.
//
// Every exported field must carry a non-empty col tag, no two fields may
// share a tag, and the tag must contain no comma. Each problem is reported
// as "FieldName: reason".
//
// Examples:
//
//	Validate(bad{}) => []string{"B: duplicate tag \"a\""}
func Validate(v any) []string {
	panic("not implemented")
}
