// Package deepcopy — Gopher Workplace challenge.
package deepcopy

import "reflect"

func deepCopy(rv reflect.Value) reflect.Value {
	switch rv.Kind() {
	case reflect.Pointer:
		if rv.IsNil() {
			return rv
		}
		out := reflect.New(rv.Type().Elem())
		out.Elem().Set(deepCopy(rv.Elem()))
		return out
	case reflect.Interface:
		if rv.IsNil() {
			return rv
		}
		out := reflect.New(rv.Type()).Elem()
		out.Set(deepCopy(rv.Elem()))
		return out
	case reflect.Slice:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Len())
		for i := 0; i < rv.Len(); i++ {
			out.Index(i).Set(deepCopy(rv.Index(i)))
		}
		return out
	case reflect.Array:
		out := reflect.New(rv.Type()).Elem()
		for i := 0; i < rv.Len(); i++ {
			out.Index(i).Set(deepCopy(rv.Index(i)))
		}
		return out
	case reflect.Map:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeMapWithSize(rv.Type(), rv.Len())
		iter := rv.MapRange()
		for iter.Next() {
			out.SetMapIndex(deepCopy(iter.Key()), deepCopy(iter.Value()))
		}
		return out
	case reflect.Struct:
		out := reflect.New(rv.Type()).Elem()
		for i := 0; i < rv.NumField(); i++ {
			if rv.Type().Field(i).IsExported() {
				out.Field(i).Set(deepCopy(rv.Field(i)))
			}
		}
		return out
	default:
		return rv
	}
}

// DeepCopy returns a copy of v that shares no mutable storage with it.
//
// Structs, slices, maps, arrays and pointers are copied recursively; scalars
// and strings are copied by value. Cycles are not part of the input.
//
// Examples:
//
//	DeepCopy(node{Tags: []string{"a"}}) => an independent node
func DeepCopy(v any) any {
	panic("not implemented")
}
