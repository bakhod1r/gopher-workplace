# Length Of Whatever You Are Handed

## Intuition

Reflection will answer "how long is this" for any type that has an answer — but only if you ask about the right kinds. The switch turns a potential panic into a second return value.

## Approach

1. Take `reflect.ValueOf(v)`.
2. Switch on the kind; for the five sized kinds return `rv.Len(), true`.
3. Default to `0, false`.

## Solution

```go
import "reflect"

// Length returns the length of v when it has one — a string, slice,
// array, map or channel — and reports false otherwise.
//
// Examples:
//
// 	Length([]int{1, 2}) => 2, true
func Length(v any) (int, bool) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return rv.Len(), true
	default:
		return 0, false
	}
}
```

## Walkthrough

`Length(nil)` gives an invalid Value whose kind is `Invalid`, which falls to the default. `Length([]int(nil))` gives a valid slice Value of length 0.

## Pitfalls

- Calling `rv.Len()` before the switch, which panics on an int.
- Forgetting that `*[3]int` is a pointer, not an array — it has no length.
