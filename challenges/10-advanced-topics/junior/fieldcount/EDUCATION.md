# How Many Fields Does This Have

## Intuition

Everything about a struct's shape is available at run time — how many fields, what they are called, and which of them the rest of the program is allowed to touch.

## Approach

1. Take `reflect.TypeOf(v)`; bail out for nil or non-struct.
2. Set `total` from `NumField`, then count the exported ones.

## Solution

```go
import "reflect"

// FieldCount returns how many fields v's struct type has in total, and
// how many of them are exported.
//
// A non-struct, or a nil interface, reports 0, 0.
//
// Examples:
//
// 	FieldCount(rec{}) => 3, 2
func FieldCount(v any) (total, exported int) {
	t := reflect.TypeOf(v)
	if t == nil || t.Kind() != reflect.Struct {
		return 0, 0
	}
	total = t.NumField()
	for i := 0; i < total; i++ {
		if t.Field(i).IsExported() {
			exported++
		}
	}
	return total, exported
}
```

## Walkthrough

`rec` has three fields; `hidden` starts with a lower-case letter, so `IsExported` is false and the exported count is 2.

## Pitfalls

- Inferring export status from the first character instead of asking; the rule involves Unicode, not just ASCII.
- Forgetting that a pointer must be dereferenced first — here the spec says not to.
