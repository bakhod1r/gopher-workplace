# Read A Struct Tag

## Intuition

Tags are just a string attached to each field. The `StructTag` type parses the conventional format for you, and `Lookup` is the version that distinguishes "absent" from "present but empty".

## Approach

1. Guard against a nil or non-struct type.
2. `FieldByName(field)`; return false if it is missing.
3. Return `f.Tag.Lookup(key)` directly — its results already match the signature.

## Solution

```go
import "reflect"

// Tag returns the value of the given key in the named field's struct tag.
//
// The second result reports whether the field exists and carries that key.
//
// Examples:
//
// 	Tag(row{}, "ID", "json") => "id", true
func Tag(v any, field, key string) (string, bool) {
	t := reflect.TypeOf(v)
	if t == nil || t.Kind() != reflect.Struct {
		return "", false
	}
	f, ok := t.FieldByName(field)
	if !ok {
		return "", false
	}
	return f.Tag.Lookup(key)
}
```

## Walkthrough

`row.ID` carries `json:"id" db:"row_id"`. Looking up "db" parses the tag and returns "row_id", true. Looking up "db" on `Name` returns "", false.

## Pitfalls

- Using `Tag.Get`, which cannot report absence.
- Forgetting that `FieldByName` also finds promoted fields from embedded structs.
