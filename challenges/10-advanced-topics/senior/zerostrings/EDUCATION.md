# Blank Every String, However Deep

## Intuition

Reflection's write rules follow the language's: you can only assign through something addressable. Starting from a pointer makes the whole tree beneath it addressable, so one recursive walk can rewrite every leaf.

## Approach

1. Verify `ptr` is a non-nil pointer to a struct.
2. Step to the struct with `Elem`.
3. Call the recursive helper.

## Solution

```go
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
// 	Redact(&record{Name: "x"}) => nil, record.Name is ""
func Redact(ptr any) error {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return ErrTarget
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return ErrTarget
	}
	redact(rv)
	return nil
}
```

## Walkthrough

From `&record{}`, `Elem` is addressable; each exported field inherits that, so the string leaves under `In`, `*Ptr` and every `List` element can be set.

## Pitfalls

- Starting from `reflect.ValueOf(v)` of a value, where nothing is settable and every write is silently skipped.
- Recursing into unexported fields, whose values cannot be set and may panic on access.
- Following a nil pointer into `Elem`, which yields an invalid Value.
