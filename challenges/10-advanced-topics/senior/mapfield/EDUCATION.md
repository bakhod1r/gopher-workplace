# Create The Map A Struct Field Needs

## Intuition

Reflection mirrors the language exactly: a nil map cannot be written to, and the fix is the same as in ordinary code — make one. The only difference is that the map's type comes from the field rather than from the source.

## Approach

1. Validate the pointer, the struct, and the `Tags` field's settability and kind.
2. Check the map's key and element kinds.
3. If the field is nil, `Set` it to `reflect.MakeMap(mt)`.
4. `SetMapIndex` with the key and value.

## Solution

```go
import (
	"errors"
	"reflect"
)

// ErrTarget is returned when ptr has no settable Tags map field.
var ErrTarget = errors.New("target must be a pointer to a struct with a settable Tags map[string]string")

// PutTag sets ptr's Tags map entry, creating the map when the field is
// nil.
//
// Writing to a nil map panics, and reflection will not create one for you.
//
// Examples:
//
// 	PutTag(&doc{}, "a", "1") => nil, doc.Tags["a"] == "1"
func PutTag(ptr any, key, val string) error {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return ErrTarget
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return ErrTarget
	}
	f := rv.FieldByName("Tags")
	if !f.IsValid() || !f.CanSet() || f.Kind() != reflect.Map {
		return ErrTarget
	}
	mt := f.Type()
	if mt.Key().Kind() != reflect.String || mt.Elem().Kind() != reflect.String {
		return ErrTarget
	}
	if f.IsNil() {
		f.Set(reflect.MakeMap(mt))
	}
	f.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
	return nil
}
```

## Walkthrough

For a fresh `doc`, `f.IsNil()` is true, so a new `map[string]string` is created and stored in the field; then the entry is written. For a doc that already has tags, the existing map is written through.

## Pitfalls

- Skipping the key/value type check, which panics inside `SetMapIndex`.
- Building the map with `reflect.MakeMap(reflect.TypeOf(map[string]string{}))` — it works here and breaks for any other map type.
