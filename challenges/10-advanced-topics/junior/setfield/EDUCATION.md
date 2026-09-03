# Write A Field Through A Pointer

## Intuition

Reflection follows Go's own rules: you cannot assign to a copy. `ValueOf` gives you a copy of whatever was in the interface, so writing requires starting from a pointer and stepping through `Elem` to the real storage.

## Approach

1. Verify `ptr` is a non-nil pointer; take `Elem()`.
2. Verify it is a struct and find the field by name.
3. Verify the field is valid, settable and of kind int, then `SetInt`.

## Solution

```go
import (
	"errors"
	"reflect"
)

// ErrNotSettable is returned when the target cannot be written.
var ErrNotSettable = errors.New("target is not a settable int field")

// SetInt sets the named int field of the struct ptr points at.
//
// Reflection can only write through a pointer: a value passed by interface
// is a copy, and the reflect package refuses to modify it.
//
// Examples:
//
// 	SetInt(&counters{}, "Hits", 3) => nil, Hits is 3
func SetInt(ptr any, field string, n int) error {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return ErrNotSettable
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return ErrNotSettable
	}
	f := rv.FieldByName(field)
	if !f.IsValid() || !f.CanSet() || f.Kind() != reflect.Int {
		return ErrNotSettable
	}
	f.SetInt(int64(n))
	return nil
}
```

## Walkthrough

`&counters{}` gives a ptr Value; `Elem()` is the addressable struct; `FieldByName("Hits")` is an addressable exported int, so `CanSet` is true and the write lands in the caller's struct.

## Pitfalls

- Skipping `CanSet` — an unexported field is valid and of the right kind, and setting it panics.
- `SetInt` on a string field panics; the kind check is not optional.
