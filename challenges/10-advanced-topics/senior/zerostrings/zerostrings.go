// Package zerostrings — Gopher Workplace challenge.
package zerostrings

import (
	"errors"
	"reflect"
)

// ErrTarget is returned when ptr is not a non-nil pointer to a struct.
var ErrTarget = errors.New("target must be a non-nil pointer to a struct")

func redact(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.String:
		if rv.CanSet() {
			rv.SetString("")
		}
	case reflect.Pointer, reflect.Interface:
		if !rv.IsNil() {
			redact(rv.Elem())
		}
	case reflect.Struct:
		rt := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			if rt.Field(i).IsExported() {
				redact(rv.Field(i))
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			redact(rv.Index(i))
		}
	}
}

// Redact sets every exported string field of the struct ptr points at to
// "", descending into nested structs and slices of structs.
//
// Unexported fields are left alone.
//
// Examples:
//
//	Redact(&record{Name: "x"}) => nil, record.Name is ""
func Redact(ptr any) error {
	panic("not implemented")
}
