// Package fielddiff — Gopher Workplace challenge.
package fielddiff

import "reflect"

func diff(av, bv reflect.Value, prefix string, out *[]string) {
	if av.Kind() == reflect.Struct {
		rt := av.Type()
		for i := 0; i < av.NumField(); i++ {
			f := rt.Field(i)
			if !f.IsExported() {
				continue
			}
			name := f.Name
			if prefix != "" {
				name = prefix + "." + name
			}
			diff(av.Field(i), bv.Field(i), name, out)
		}
		return
	}
	if !av.Equal(bv) {
		*out = append(*out, prefix)
	}
}

// Diff returns the dotted paths of the exported fields where a and b
// differ, in declaration order.
//
// a and b must have the same type; otherwise the result is nil. Nested
// structs contribute dotted paths.
//
// Examples:
//
//	Diff(cfg{A: 1}, cfg{A: 2}) => []string{"A"}
func Diff(a, b any) []string {
	panic("not implemented")
}
