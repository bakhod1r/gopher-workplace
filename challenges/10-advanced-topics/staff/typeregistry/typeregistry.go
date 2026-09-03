// Package typeregistry — Gopher Workplace challenge.
package typeregistry

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
//	New("job") => *job, nil
func New(name string) (any, error) {
	panic("not implemented")
}
