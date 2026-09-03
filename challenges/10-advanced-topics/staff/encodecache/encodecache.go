// Package encodecache — Gopher Workplace challenge.
package encodecache

import (
	"errors"
	"reflect"
	"sync"
)

// ErrKind is returned when v is not a struct.
var ErrKind = errors.New("v must be a struct")

// layouts caches each struct type's exported string field indices.
var layouts sync.Map // reflect.Type -> []fieldRef

type fieldRef struct {
	name  string
	index int
}

// layoutOf resolves and caches the encodable fields of t.
func layoutOf(t reflect.Type) []fieldRef {
	if v, ok := layouts.Load(t); ok {
		return v.([]fieldRef)
	}
	refs := make([]fieldRef, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.IsExported() && f.Type.Kind() == reflect.String {
			refs = append(refs, fieldRef{name: f.Name, index: i})
		}
	}
	actual, _ := layouts.LoadOrStore(t, refs)
	return actual.([]fieldRef)
}

// Encode appends "name=value" for each exported string field of v to dst,
// separated by '&', and returns the extended slice.
//
// The per-type field list is resolved once and cached, so repeated
// encodings of a known type cost a map lookup and some appends.
//
// Examples:
//
//	Encode(nil, user{Name: "a"}) => []byte("Name=a")
func Encode(dst []byte, v any) ([]byte, error) {
	panic("not implemented")
}
