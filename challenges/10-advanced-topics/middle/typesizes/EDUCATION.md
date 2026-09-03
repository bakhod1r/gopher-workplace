# Size And Alignment Of A Run-Time Type

## Intuition

`unsafe.Sizeof` is a compile-time constant, so it cannot answer for a type you only meet at run time. `reflect.Type` carries the same numbers, computed by the same rules.

## Approach

1. Take `reflect.TypeOf(v)`; return false if it is nil.
2. Return `t.Size()` and `uintptr(t.Align())`.

## Solution

```go
import "reflect"

// Sizes returns the size and alignment of v's dynamic type.
//
// A nil interface has no type, so it reports false.
//
// Examples:
//
// 	Sizes(int64(0)) => 8, 8, true
func Sizes(v any) (size, align uintptr, ok bool) {
	t := reflect.TypeOf(v)
	if t == nil {
		return 0, 0, false
	}
	return t.Size(), uintptr(t.Align()), true
}
```

## Walkthrough

`Sizes(wide{})` reports 16 and 8: one byte, seven bytes of padding, then the int64. The alignment is the widest field's.

## Pitfalls

- Expecting a slice's size to include its elements.
- `reflect.TypeOf(nil)` returns nil, and calling `Size` on it panics.
