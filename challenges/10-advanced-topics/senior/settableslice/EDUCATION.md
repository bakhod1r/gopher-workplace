# Write Into A Slice Through Reflection

## Intuition

A `reflect.Value` is a handle onto storage. Assigning to the Go variable holding the handle changes only the variable — the storage is written through `Set` methods, and only when the handle is addressable.

## Approach

1. Take each element's Value with `rv.Index(i)`.
2. Write with `e.SetInt(e.Int() * 2)`.

## Solution

```go
import (
	"errors"
	"reflect"
)

// ErrShape is returned when slice is not a slice of ints.
var ErrShape = errors.New("argument must be a slice of ints")

// Double multiplies every element of the int slice in place.
//
// reflect.ValueOf gives a copy of the interface's contents, but a slice's
// elements live in the shared backing array — which is exactly why the
// elements are settable and the slice header is not.
//
// Examples:
//
// 	s := []int{1, 2}; Double(s) => s is [2 4]
func Double(slice any) error {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice || rv.Type().Elem().Kind() != reflect.Int {
		return ErrShape
	}
	for i := 0; i < rv.Len(); i++ {
		e := rv.Index(i)
		e.SetInt(e.Int() * 2)
	}
	return nil
}
```

## Walkthrough

`rv.Index(0)` is an addressable handle onto `s[0]`. `SetInt` writes 2 into it. The buggy line built a new, unaddressable Value and dropped it.

## Pitfalls

- Expecting `reflect.ValueOf(s)` to be unsettable to make the elements unsettable too — the header is a copy, the array is not.
- `e.Set(reflect.ValueOf(v*2))` also works; the point is that a `Set` call is what is missing.
