# Fill A Struct From A String Map

## Intuition

Decoding is a table walk. The struct type is the table — each row has a name, a kind and a tag — and the map is the data. Reflection lets you write the walk once instead of once per config struct.

## Approach

1. Validate `dst` and step to the struct with `Elem`.
2. For each field, read the `env` tag and skip the exclusions.
3. Look the key up in `src`; skip if absent.
4. Switch on the field's kind, parse, and set. Wrap parse failures with the field name.

## Solution

```go
import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ErrTarget is returned when dst is not a non-nil pointer to a struct.
var ErrTarget = errors.New("dst must be a non-nil pointer to a struct")

// Decode fills dst's fields from src, matching by the field's `env` tag.
//
// Supported field kinds are string, int and bool. Fields without an env
// tag, unexported fields, and keys missing from src are left alone.
//
// Examples:
//
// 	Decode(map[string]string{"PORT": "80"}, &cfg) => cfg.Port == 80
func Decode(src map[string]string, dst any) error {
	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return ErrTarget
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return ErrTarget
	}
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		key, ok := f.Tag.Lookup("env")
		if !ok || key == "" || key == "-" || !f.IsExported() {
			continue
		}
		s, ok := src[key]
		if !ok {
			continue
		}
		fv := rv.Field(i)
		switch fv.Kind() {
		case reflect.String:
			fv.SetString(s)
		case reflect.Int:
			n, err := strconv.Atoi(s)
			if err != nil {
				return fmt.Errorf("field %s: %w", f.Name, err)
			}
			fv.SetInt(int64(n))
		case reflect.Bool:
			b, err := strconv.ParseBool(s)
			if err != nil {
				return fmt.Errorf("field %s: %w", f.Name, err)
			}
			fv.SetBool(b)
		default:
			return fmt.Errorf("field %s: unsupported kind %s", f.Name, fv.Kind())
		}
	}
	return nil
}
```

## Walkthrough

With `PORT` set to "8080", the loop reaches `Port`, reads tag "PORT", finds the value, sees kind int, parses 8080 and calls `SetInt`. `Ignored` has no tag and is never touched.

## Pitfalls

- Checking `IsExported` after calling `Set` — the panic comes first.
- Treating `env:"-"` as a key named "-".
- Returning early on the first missing key instead of skipping it.
