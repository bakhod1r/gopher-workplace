# Report Which Fields Differ

## Intuition

Two values of the same type have the same shape, so the walk visits identical positions in both. Only the leaves need comparing, and the path to each leaf is whatever names you passed on the way down.

## Approach

1. Reject invalid Values and mismatched types.
2. Call the helper with an empty prefix and a pointer to the result slice.
3. Return the accumulated paths.

## Solution

```go
import "reflect"

func diff(av, bv reflect.Value, prefix string, out *[]string) {
	if av.Kind() == reflect.Struct {
		rt := av.Type()
		for i := 0; i < av.NumField(); i++ {
			f := rt.Field(i)
			if !f.IsExported() {
				continue
			}
			name := f.Name
			if prefix != "" {
				name = prefix + "." + name
			}
			diff(av.Field(i), bv.Field(i), name, out)
		}
		return
	}
	if !av.Equal(bv) {
		*out = append(*out, prefix)
	}
}

// Diff returns the dotted paths of the exported fields where a and b
// differ, in declaration order.
//
// a and b must have the same type; otherwise the result is nil. Nested
// structs contribute dotted paths.
//
// Examples:
//
// 	Diff(cfg{A: 1}, cfg{A: 2}) => []string{"A"}
func Diff(a, b any) []string {
	av, bv := reflect.ValueOf(a), reflect.ValueOf(b)
	if !av.IsValid() || !bv.IsValid() || av.Type() != bv.Type() {
		return nil
	}
	var out []string
	diff(av, bv, "", &out)
	return out
}
```

## Walkthrough

For `settings`, the walk visits Name, Retry, then descends into Limits and visits Limits.Soft and Limits.Hard. `hidden` is skipped. Only the differing leaves are appended, in that order.

## Pitfalls

- Comparing structs at the top level with `Equal` — you learn that they differ, not where.
- Using `reflect.DeepEqual` on the leaves, which boxes both sides on every comparison.
- Descending into unexported fields, which `Equal` refuses to compare.
