# Reject A Bad Schema Before It Ships

## Intuition

A tag-driven mapping is a schema the compiler does not check. Walking the type once at start-up turns a silent production bug into a startup failure.

## Approach

1. Reject non-structs.
2. For each exported field: look the tag up, then check empty, comma and duplicate in that order.
3. Record the tag's owner in `seen` only when the field is otherwise valid.

## Solution

```go
import "reflect"

// Validate returns the problems with v's `col` tags, in field order.
//
// Every exported field must carry a non-empty col tag, no two fields may
// share a tag, and the tag must contain no comma. Each problem is reported
// as "FieldName: reason".
//
// Examples:
//
// 	Validate(bad{}) => []string{"B: duplicate tag \"a\""}
func Validate(v any) []string {
	rt := reflect.TypeOf(v)
	if rt == nil || rt.Kind() != reflect.Struct {
		return []string{"not a struct"}
	}
	var out []string
	seen := make(map[string]string, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if !f.IsExported() {
			continue
		}
		tag, ok := f.Tag.Lookup("col")
		if !ok || tag == "" {
			out = append(out, f.Name+": missing col tag")
			continue
		}
		bad := false
		for j := 0; j < len(tag); j++ {
			if tag[j] == ',' {
				bad = true
				break
			}
		}
		if bad {
			out = append(out, f.Name+": tag contains a comma")
			continue
		}
		if prev, dup := seen[tag]; dup {
			out = append(out, f.Name+": duplicate tag of "+prev)
			continue
		}
		seen[tag] = f.Name
	}
	return out
}
```

## Walkthrough

For `multi`: A claims "x"; B has no tag and is reported; C's "x" is already claimed by A, so it is reported as a duplicate of A.

## Pitfalls

- Recording a tag in `seen` before validating it, so a later field duplicates an invalid tag.
- Using `Tag.Get`, which cannot tell an empty tag from a missing one.
