# Reflect Only When You Have To

## Intuition

A type switch is a handful of comparisons the compiler lays out; reflection is a run-time inspection. Handling the common types first means the general machinery only runs for the cases that need it.

## Approach

1. Type-switch the concrete types, appending directly.
2. Otherwise take `reflect.ValueOf(v)` and switch on the kind.
3. Use the typed accessors — `String`, `Int`, `Uint`, `Bool`.
4. Default to `?`.

## Solution

```go
import (
	"reflect"
	"strconv"
)

// Render appends v's text form to dst.
//
// The common types are handled by a type switch, which costs nothing;
// everything else falls back to reflection. The fast path must not
// allocate.
//
// Examples:
//
// 	Render(nil, 42) => []byte("42")
func Render(dst []byte, v any) []byte {
	switch x := v.(type) {
	case nil:
		return append(dst, "<nil>"...)
	case string:
		return append(dst, x...)
	case int:
		return strconv.AppendInt(dst, int64(x), 10)
	case int64:
		return strconv.AppendInt(dst, x, 10)
	case bool:
		return strconv.AppendBool(dst, x)
	case []byte:
		return append(dst, x...)
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String:
		return append(dst, rv.String()...)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.AppendInt(dst, rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.AppendUint(dst, rv.Uint(), 10)
	case reflect.Bool:
		return strconv.AppendBool(dst, rv.Bool())
	default:
		return append(dst, '?')
	}
}
```

## Walkthrough

`Render(dst, 42)` matches `case int` and appends two digits. `Render(dst, myInt(5))` falls through the switch, reaches the fallback, and `rv.Int()` reads it without boxing.

## Pitfalls

- `fmt.Append` in the fallback, which is correct and allocates for the boxing.
- Forgetting that `case nil` in a type switch matches a nil interface, not a typed nil.
- Using `rv.Interface()` in the fallback, which allocates and undoes the point.
