# Is This Value The Zero Value

## Intuition

"Zero" is a per-type notion — 0, "", nil, an all-zero struct. Reflection knows the type, so it can answer the question generically, but only once you have handled the case where there is no type at all.

## Approach

1. `reflect.ValueOf(v)`.
2. Return true when the Value is invalid — a nil interface holds nothing.
3. Otherwise return `rv.IsZero()`.

## Solution

```go
import "reflect"

// IsZero reports whether v holds the zero value for its type.
//
// A nil interface counts as zero: there is nothing in it.
//
// Examples:
//
// 	IsZero(0) => true
func IsZero(v any) bool {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return true
	}
	return rv.IsZero()
}
```

## Walkthrough

`IsZero((*pair)(nil))` sees a valid Value of kind ptr whose value is nil, so `IsZero` is true. `IsZero(&pair{})` sees a non-nil pointer, so it is false even though what it points at is zero.

## Pitfalls

- `v == nil` on the interface misses a typed nil pointer, which is not `nil` as an interface.
- Calling `IsZero` on the invalid Value from `reflect.ValueOf(nil)` panics.
