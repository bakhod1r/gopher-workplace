# Grow A Slice Whose Type You Do Not Know

## Intuition

Reflection mirrors the language: `append` returns a new slice header, so the reflective version does too. The extra step is storing it back through the addressable Value you got from `Elem`.

## Approach

1. Validate the pointer, step to the slice with `Elem`, and reject a negative `n`.
2. Take the element type from `rv.Type().Elem()`.
3. Append `reflect.Zero(elem)` `n` times, storing each result with `rv.Set`.

## Solution

```go
import (
	"errors"
	"reflect"
)

// ErrTarget is returned when slicePtr is not a non-nil pointer to a slice.
var ErrTarget = errors.New("target must be a non-nil pointer to a slice")

// AppendZero appends n zero values to the slice that slicePtr points at.
//
// The element type comes from the slice itself, so one implementation
// serves every slice type.
//
// Examples:
//
// 	s := []int{1}; AppendZero(&s, 2) => s is [1 0 0]
func AppendZero(slicePtr any, n int) error {
	rv := reflect.ValueOf(slicePtr)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return ErrTarget
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Slice {
		return ErrTarget
	}
	if n < 0 {
		return ErrTarget
	}
	elem := rv.Type().Elem()
	for i := 0; i < n; i++ {
		rv.Set(reflect.Append(rv, reflect.Zero(elem)))
	}
	return nil
}
```

## Walkthrough

`&s` where `s` is `[]int{1}` gives an addressable slice Value. Two appends of `reflect.Zero(int)` produce [1 0] then [1 0 0], each stored back into the caller's slice header.

## Pitfalls

- Discarding `Append`'s result — the caller's slice never changes.
- Building the zero value as `reflect.ValueOf(0)`, which only works for int.
