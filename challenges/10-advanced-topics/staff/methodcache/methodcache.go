// Package methodcache — Gopher Workplace challenge.
package methodcache

import (
	"errors"
	"reflect"
	"sync"
)

// ErrMethod is returned when the method is missing or has the wrong shape.
var ErrMethod = errors.New("no such method with signature func(int) int")

type methodKey struct {
	t    reflect.Type
	name string
}

// methods caches the resolved method index, or -1 when the lookup failed.
var methods sync.Map // methodKey -> int

// methodIndex resolves the method's index on t, caching the answer.
func methodIndex(t reflect.Type, name string) int {
	k := methodKey{t, name}
	if v, ok := methods.Load(k); ok {
		return v.(int)
	}
	idx := -1
	if m, ok := t.MethodByName(name); ok {
		mt := m.Type
		// mt includes the receiver as parameter 0.
		if mt.NumIn() == 2 && mt.In(1).Kind() == reflect.Int &&
			mt.NumOut() == 1 && mt.Out(0).Kind() == reflect.Int {
			idx = m.Index
		}
	}
	actual, _ := methods.LoadOrStore(k, idx)
	return actual.(int)
}

// CallNamed calls the named method on v with one int argument and
// returns its single int result.
//
// Method lookup by name is a search; the resolved index is cached per
// (type, name) so repeated calls cost a map lookup.
//
// Examples:
//
//	CallNamed(adder{2}, "Add", 3) => 5, nil
func CallNamed(v any, method string, arg int) (int, error) {
	panic("not implemented")
}
