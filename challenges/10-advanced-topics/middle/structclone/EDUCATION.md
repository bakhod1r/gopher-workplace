# Copy A Struct Without Knowing Its Type

## Intuition

Reflection can build a value of any type it is shown. `New` allocates one, `Elem` gives the addressable value inside, and `Set` performs exactly the assignment the language would.

## Approach

1. Reject a nil interface and non-struct kinds.
2. `reflect.New(rv.Type()).Elem()` for a fresh addressable value.
3. `Set` the original into it and return `Interface()`.

## Solution

```go
import "reflect"

// Clone returns a copy of the struct v, as a value of the same type.
//
// The copy is shallow: fields are assigned, so slices and maps inside it
// still share their storage with v.
//
// Examples:
//
// 	Clone(pt{1, 2}) => pt{1, 2}, a distinct value
func Clone(v any) any {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return nil
	}
	out := reflect.New(rv.Type()).Elem()
	out.Set(rv)
	return out.Interface()
}
```

## Walkthrough

`Clone(pt{1,2})` allocates a new `pt`, assigns the original's fields, and boxes the result. Assigning to `out.X` afterwards cannot reach `in`, but writing through `out.Tags` can — that is what shallow means.

## Pitfalls

- Returning `out` (a `reflect.Value`) instead of `out.Interface()`.
- Promising a deep copy; slices and maps are still shared.
