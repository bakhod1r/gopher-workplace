// Package walkfields — Gopher Workplace challenge.
package walkfields

import "reflect"

func deepSum(rv reflect.Value) int64 {
	if !rv.IsValid() {
		return 0
	}
	switch rv.Kind() {
	case reflect.Int:
		return rv.Int()
	case reflect.Pointer, reflect.Interface:
		if rv.IsNil() {
			return 0
		}
		return deepSum(rv.Elem())
	case reflect.Struct:
		rt := rv.Type()
		var total int64
		for i := 0; i < rv.NumField(); i++ {
			if !rt.Field(i).IsExported() {
				continue
			}
			total += deepSum(rv.Field(i))
		}
		return total
	case reflect.Slice, reflect.Array:
		var total int64
		for i := 0; i < rv.Len(); i++ {
			total += deepSum(rv.Index(i))
		}
		return total
	default:
		return 0
	}
}

// DeepSum totals every exported int field in v, descending into nested
// structs, slices of structs and pointers.
//
// A nil pointer contributes nothing. Cycles are not part of the input.
//
// Examples:
//
//	DeepSum(outer{N: 1, In: inner{M: 2}}) => 3
func DeepSum(v any) int64 {
	panic("not implemented")
}
