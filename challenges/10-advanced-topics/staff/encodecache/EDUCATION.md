# Encode Any Struct, Resolve It Once

## Intuition

Reflection is expensive where it inspects and cheap where it accesses. Resolving the layout once moves all the inspection to the first call, leaving a loop of appends that costs about what hand-written code would.

## Approach

1. Reject a non-struct.
2. Fetch the cached field references for the type.
3. For each, append the separator, the name, `=`, and the field's string value.

## Solution

```go
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
// 	Encode(nil, user{Name: "a"}) => []byte("Name=a")
func Encode(dst []byte, v any) ([]byte, error) {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return dst, ErrKind
	}
	refs := layoutOf(rv.Type())
	for i, r := range refs {
		if i > 0 {
			dst = append(dst, '&')
		}
		dst = append(dst, r.name...)
		dst = append(dst, '=')
		dst = append(dst, rv.Field(r.index).String()...)
	}
	return dst, nil
}
```

## Walkthrough

The first encode of `user` walks four fields and stores two references. Every later one — from any goroutine — does a `sync.Map` load and four appends into the caller's buffer.

## Pitfalls

- Ranging the struct's fields instead of the cached layout, which restores the per-call walk.
- `rv.Field(i).Interface().(string)`, which boxes and allocates on every field.
- Caching a `reflect.Value`; Values are bound to one variable, indices are not.
