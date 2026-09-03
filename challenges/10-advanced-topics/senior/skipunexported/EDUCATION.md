# Sum The Fields You Are Allowed To Read

## Intuition

Reflection enforces the language's visibility rules: it will show you an unexported field but refuse to hand it out as an `any`. And a field of kind int holds an `int`, not an `int64` — the type assertion was never going to succeed.

## Approach

1. Consult `rt.Field(i).IsExported()` and skip unexported fields.
2. Read the value with `f.Int()`, which returns int64 for every signed integer kind.

## Solution

```go
import "reflect"

// SumInts returns the total of v's exported int fields.
//
// Unexported fields can be read as reflect Values but not converted back
// through Interface, and reaching for them panics.
//
// Examples:
//
// 	SumInts(mix{A: 1, b: 2}) => 1
func SumInts(v any) int64 {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return 0
	}
	rt := rv.Type()
	var total int64
	for i := 0; i < rv.NumField(); i++ {
		if !rt.Field(i).IsExported() {
			continue
		}
		f := rv.Field(i)
		if f.Kind() != reflect.Int {
			continue
		}
		total += f.Int()
	}
	return total
}
```

## Walkthrough

For `mix{A:1, B:2, hidden:100}`: A and B are exported ints and contribute 3; `hidden` is skipped before anything can panic; `Name` fails the kind check.

## Pitfalls

- `f.Interface()` on an unexported field panics with "cannot return value obtained from unexported field".
- `f.Interface().(int64)` panics even on an exported int field — the dynamic type is `int`.
