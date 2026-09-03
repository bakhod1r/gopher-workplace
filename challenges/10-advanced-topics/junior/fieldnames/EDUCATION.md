# List A Struct's Exported Fields

## Intuition

A struct type carries its full field list at run time: names, types, tags and export status. Reflection just reads that table.

## Approach

1. Take `reflect.TypeOf(v)`; return nil if it is nil or not a struct.
2. Loop over `NumField()`, appending each exported field's `Name`.

## Solution

```go
import "reflect"

// FieldNames returns the names of v's exported fields, in declaration
// order.
//
// A non-struct, or a nil interface, yields nil.
//
// Examples:
//
// 	FieldNames(struct{ A, b int }{}) => []string{"A"}
func FieldNames(v any) []string {
	t := reflect.TypeOf(v)
	if t == nil || t.Kind() != reflect.Struct {
		return nil
	}
	var out []string
	for i := 0; i < t.NumField(); i++ {
		if f := t.Field(i); f.IsExported() {
			out = append(out, f.Name)
		}
	}
	return out
}
```

## Walkthrough

`user` has three fields; `admin` starts with a lower-case letter, so `IsExported` is false and it is skipped, leaving [Name Age].

## Pitfalls

- Calling `NumField` before checking the kind — that is a panic, not an error.
- Dereferencing pointers; the spec here says a pointer yields nil.
