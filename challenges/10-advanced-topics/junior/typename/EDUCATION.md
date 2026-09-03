# What Type Is In This Interface

## Intuition

An interface value carries a type word. Reflection reads it — and when the interface is nil there is nothing to read, which is why the nil check comes first.

## Approach

1. `reflect.TypeOf(v)`.
2. Return `"<nil>"` when it is nil.
3. Return `t.String()`.

## Solution

```go
import "reflect"

// TypeName returns the name of v's dynamic type.
//
// A nil interface holds no type at all, so it reports "<nil>".
//
// Examples:
//
// 	TypeName(3) => "int"
func TypeName(v any) string {
	t := reflect.TypeOf(v)
	if t == nil {
		return "<nil>"
	}
	return t.String()
}
```

## Walkthrough

`TypeName(named(1))` reports the declared name qualified by its package, while `TypeName(alias(1))` reports "int" because an alias declares no new type.

## Pitfalls

- Calling `t.String()` without the nil check — that is a nil-pointer panic.
- Expecting `[]byte` to print as "[]byte"; `byte` is an alias for `uint8`.
