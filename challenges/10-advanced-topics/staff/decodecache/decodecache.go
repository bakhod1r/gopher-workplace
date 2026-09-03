// Package decodecache — Gopher Workplace challenge.
package decodecache

import (
	"errors"
	"reflect"
	"sync"
)

// ErrTarget is returned when dst is not a non-nil pointer to a struct.
var ErrTarget = errors.New("dst must be a non-nil pointer to a struct")

// layouts caches the tag-to-field-index map for each struct type.
var layouts sync.Map // reflect.Type -> map[string]int

// layoutOf returns the tag-to-index map for t, computing it at most once
// per type as far as any caller can observe.
func layoutOf(t reflect.Type) map[string]int {
	if v, ok := layouts.Load(t); ok {
		return v.(map[string]int)
	}
	m := make(map[string]int, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() || f.Type.Kind() != reflect.String {
			continue
		}
		if key, ok := f.Tag.Lookup("env"); ok && key != "" && key != "-" {
			m[key] = i
		}
	}
	actual, _ := layouts.LoadOrStore(t, m)
	return actual.(map[string]int)
}

// Decode fills dst's string fields from src by their `env` tag, caching
// each struct type's tag-to-index layout.
//
// The cache is shared by concurrent callers, so it must be safe under
// parallel use — and must not resolve the layout twice for one type.
//
// Examples:
//
//	Decode(map[string]string{"H": "x"}, &cfg) => cfg.H == "x"
func Decode(src map[string]string, dst any) error {
	panic("not implemented")
}
