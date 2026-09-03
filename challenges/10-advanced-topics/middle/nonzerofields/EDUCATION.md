# Which Fields Were Actually Set

## Intuition

Reflection lets one function answer "what changed" for every struct in the codebase, and `IsZero` handles each field's notion of empty without a type switch.

## Approach

1. Reject non-structs.
2. Loop the fields, skipping unexported ones.
3. Append the name when `rv.Field(i).IsZero()` is false.

## Solution

```go
import "reflect"

// NonZero returns the names of v's exported fields that hold something
// other than their zero value, in declaration order.
//
// A non-struct, or a nil interface, yields nil.
//
// Examples:
//
// 	NonZero(patch{Name: "x"}) => []string{"Name"}
func NonZero(v any) []string {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return nil
	}
	rt := rv.Type()
	var out []string
	for i := 0; i < rv.NumField(); i++ {
		if !rt.Field(i).IsExported() {
			continue
		}
		if !rv.Field(i).IsZero() {
			out = append(out, rt.Field(i).Name)
		}
	}
	return out
}
```

## Walkthrough

`patch{Name:"x", Count:3}` has two non-zero fields; `Active` is false and `Tags` is nil, both of which are their types' zero values.

## Pitfalls

- Comparing against a fresh zero struct field by field, which needs the type at compile time.
- Treating an empty slice as absent — `IsZero` deliberately does not.
