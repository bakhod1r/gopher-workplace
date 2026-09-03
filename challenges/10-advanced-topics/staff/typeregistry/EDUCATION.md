# Build A Value From A Name

## Intuition

A type value is a run-time handle to everything the compiler knew about a type — enough to allocate one. That turns a forty-case switch into a map from name to type plus a single `New`.

## Approach

1. `lookup(name)`; return `ErrUnknown` on a nil type.
2. `reflect.New(t)` for a pointer to a fresh zero value.
3. Return `Interface()` so the caller can type-assert it.

## Solution

```go
import (
	"errors"
	"reflect"
	"sync"
)

// ErrUnknown is returned when no type is registered under the name.
var ErrUnknown = errors.New("unknown type name")

// registry maps a name to the registered type.
var registry sync.Map // string -> reflect.Type

// Register records a prototype under name. It is safe for concurrent use.
func Register(name string, prototype any) {
	registry.Store(name, reflect.TypeOf(prototype))
}

// lookup returns the registered type, or nil.
func lookup(name string) reflect.Type {
	v, ok := registry.Load(name)
	if !ok {
		return nil
	}
	t, _ := v.(reflect.Type)
	return t
}

// New returns a freshly allocated pointer to the type registered under
// name.
//
// The registry is written once at init and read by many goroutines, so the
// lookup must be safe without serialising every construction.
//
// Examples:
//
// 	New("job") => *job, nil
func New(name string) (any, error) {
	t := lookup(name)
	if t == nil {
		return nil, ErrUnknown
	}
	return reflect.New(t).Interface(), nil
}
```

## Walkthrough

`New("job")` finds `reflect.TypeOf(job{})`, allocates a zero `job`, and boxes the `*job` pointer. The caller's `v.(*job)` succeeds because the dynamic type survived the round trip.

## Pitfalls

- Returning `reflect.New(t).Elem().Interface()`, which boxes a copy and gives the caller nothing to write through.
- Caching a constructed value instead of the type, which would hand every caller the same struct.
- Registering a pointer prototype, which makes `New` return `**T`.
