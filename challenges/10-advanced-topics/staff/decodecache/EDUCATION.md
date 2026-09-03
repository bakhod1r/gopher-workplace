# Resolve The Layout Once Per Type, Safely

## Intuition

Reflection's expensive part is asking the type questions, and the answers never change. Cache the answer per type; the per-call work then drops to a map lookup and a few field writes.

## Approach

1. Validate `dst` and step to the struct.
2. Fetch the layout for `rv.Type()` from the cache.
3. For each cached tag-to-index pair, set the field when `src` has the key.

## Solution

```go
import (
	"errors"
	"reflect"
	"sync"
)

// ErrTarget is returned when dst is not a non-nil pointer to a struct.
var ErrTarget = errors.New("dst must be a non-nil pointer to a struct")

// layouts caches the tag-to-field-index map for each struct type.
var layouts sync.Map // reflect.Type -> map[string]int

// layoutOf returns the tag-to-index map for t, computing it at most once
// per type as far as any caller can observe.
func layoutOf(t reflect.Type) map[string]int {
	if v, ok := layouts.Load(t); ok {
		return v.(map[string]int)
	}
	m := make(map[string]int, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() || f.Type.Kind() != reflect.String {
			continue
		}
		if key, ok := f.Tag.Lookup("env"); ok && key != "" && key != "-" {
			m[key] = i
		}
	}
	actual, _ := layouts.LoadOrStore(t, m)
	return actual.(map[string]int)
}

// Decode fills dst's string fields from src by their `env` tag, caching
// each struct type's tag-to-index layout.
//
// The cache is shared by concurrent callers, so it must be safe under
// parallel use — and must not resolve the layout twice for one type.
//
// Examples:
//
// 	Decode(map[string]string{"H": "x"}, &cfg) => cfg.H == "x"
func Decode(src map[string]string, dst any) error {
	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return ErrTarget
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return ErrTarget
	}
	layout := layoutOf(rv.Type())
	for key, i := range layout {
		if s, ok := src[key]; ok {
			rv.Field(i).SetString(s)
		}
	}
	return nil
}
```

## Walkthrough

The first `Decode` of `cfg` walks five fields and stores a two-entry map. Every later call — from any goroutine — does one `sync.Map` load and up to two `SetString` calls.

## Pitfalls

- Ranging the struct's fields instead of the cached layout, which reintroduces the walk.
- Caching a `reflect.Value` instead of an index; Values are bound to a particular variable.
- Guarding a plain map with no lock — the concurrency test is there to catch it.
