# The Interface That Is Not Nil

## Intuition

An interface value carries a type word and a data word. A nil `*T` fills in the type word, so the interface is not nil even though the pointer inside is. Reflection is the only way to look inside and ask about the data word alone.

## Approach

1. Return true immediately for an untyped nil.
2. Switch on the kind and call `IsNil` only for the kinds that can be nil.
3. Return false for everything else.

## Solution

```go
import "reflect"

// IsNilValue reports whether v is nil or holds a nil pointer, map,
// slice, channel, function or interface.
//
// An interface holding a typed nil pointer is not == nil, which is the trap
// this function exists to close.
//
// Examples:
//
// 	var p *int; IsNilValue(p) => true
func IsNilValue(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Chan,
		reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return rv.IsNil()
	default:
		return false
	}
}
```

## Walkthrough

`var e error = (*myErr)(nil)` has type `*myErr` and value nil. `e == nil` is false; `reflect.ValueOf(e).IsNil()` is true, which is the answer the caller wanted.

## Pitfalls

- Calling `IsNil` without the kind switch — that panics on an int.
- Fixing the symptom at the call site instead of never assigning a typed nil to an interface in the first place.
