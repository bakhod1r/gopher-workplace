# Describe A Struct's Shape

## Intuition

A struct type is a table the runtime carries: names, types, tags, export status. Rendering it is a loop, and the result stays correct as the struct changes.

## Approach

1. Reject a nil or non-struct type.
2. Loop the fields, skipping unexported ones.
3. Append `Name + ":" + Type.Kind().String()`.

## Solution

```go
import "reflect"

// FieldKinds returns "Name:kind" for each exported field of v, in
// declaration order.
//
// A non-struct, or a nil interface, yields nil.
//
// Examples:
//
// 	FieldKinds(row{}) => []string{"ID:int", "Name:string"}
func FieldKinds(v any) []string {
	t := reflect.TypeOf(v)
	if t == nil || t.Kind() != reflect.Struct {
		return nil
	}
	var out []string
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}
		out = append(out, f.Name+":"+f.Type.Kind().String())
	}
	return out
}
```

## Walkthrough

`row` has five fields; `hidden` is skipped, and `Tags` reports kind slice rather than the type string "[]string".

## Pitfalls

- Using `f.Type.String()`, which gives the type name instead of the kind.
- Calling `NumField` before checking the kind.
