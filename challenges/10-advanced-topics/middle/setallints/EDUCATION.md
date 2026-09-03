# Write Every Int Field At Once

## Intuition

Once you have an addressable struct, each field is either writable or not, and reflection will tell you which. That turns "reset all counters" into a loop the struct definition drives.

## Approach

1. Validate the pointer and step to the struct.
2. For each field, write it when the kind is int and `CanSet` is true.
3. Return the count.

## Solution

```go
import (
	"errors"
	"reflect"
)

// ErrTarget is returned when ptr is not a non-nil pointer to a struct.
var ErrTarget = errors.New("target must be a non-nil pointer to a struct")

// SetAllInts sets every settable int field of the struct ptr points at to
// v, and reports how many fields it wrote.
//
// Unexported fields and fields of other kinds are skipped.
//
// Examples:
//
// 	SetAllInts(&rec{}, 7) => 2, nil for a struct with two int fields
func SetAllInts(ptr any, v int) (int, error) {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return 0, ErrTarget
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return 0, ErrTarget
	}
	n := 0
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if f.Kind() == reflect.Int && f.CanSet() {
			f.SetInt(int64(v))
			n++
		}
	}
	return n, nil
}
```

## Walkthrough

`rec` has five fields; `A` and `B` pass both checks, `Name` and `Ratio` fail the kind check, and `hidden` fails `CanSet`.

## Pitfalls

- Checking `IsExported` but not the kind, which panics on the string field.
- Using `reflect.ValueOf(ptr).Field(i)` without `Elem`, which is not a struct at all.
