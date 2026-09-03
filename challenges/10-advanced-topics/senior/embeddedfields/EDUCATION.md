# Fields That Came From Somewhere Else

## Intuition

Embedding is a naming rule, not a new kind of field. Reflection shows it as an ordinary field with the `Anonymous` flag set — so "descend or not" is one boolean, and everything else is an ordinary walk.

## Approach

1. Reject non-structs.
2. For each exported field: if it is anonymous and a struct, recurse and prefix with its name.
3. Otherwise append the field's name as a leaf.

## Solution

```go
import "reflect"

// Paths returns the dotted path of every exported leaf field of v,
// descending through embedded structs.
//
// An embedded struct contributes its fields' paths under its own name;
// named struct fields are not descended into.
//
// Examples:
//
// 	Paths(User{}) => []string{"Base.ID", "Name"}
func Paths(v any) []string {
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
		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			for _, sub := range Paths(reflect.New(f.Type).Elem().Interface()) {
				out = append(out, f.Name+"."+sub)
			}
			continue
		}
		out = append(out, f.Name)
	}
	return out
}
```

## Walkthrough

`user` embeds `Base`, so the walk recurses and prefixes, giving `Base.ID`. `Extra` is a named struct field, so it stops there.

## Pitfalls

- Treating every struct field as embedded, which flattens `Extra` too.
- Using `FieldByName` for promoted fields, which finds them without telling you where they came from.
