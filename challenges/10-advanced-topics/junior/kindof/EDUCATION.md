# Ask A Value What It Is

## Intuition

Reflection starts by opening the interface: `ValueOf` gives you a handle you can ask questions of. `Kind` is the coarsest question — what shape is this, regardless of what it is called.

## Approach

1. `reflect.ValueOf(v)`.
2. `.Kind()` for the shape, `.String()` for its name.

## Solution

```go
import "reflect"

// KindName returns the name of v's underlying kind: "int", "slice",
// "struct" and so on.
//
// A nil interface has no type at all, so it reports "invalid".
//
// Examples:
//
// 	KindName(3) => "int"
func KindName(v any) string {
	return reflect.ValueOf(v).Kind().String()
}
```

## Walkthrough

`myInt(1)` has type name `myInt` and kind `int`. `KindName` reports the kind, so both `1` and `myInt(1)` give "int".

## Pitfalls

- `reflect.TypeOf(nil)` returns nil and panics when you call `Kind` on it; `ValueOf` gives a usable zero Value instead.
- Confusing `Kind().String()` with `Type().String()` — the latter gives "main.myInt".
