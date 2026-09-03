# Call A Method By Name, Once Resolved

## Intuition

Method resolution is a property of the type, so it is cacheable exactly like field layout. The subtlety is the receiver: it occupies parameter 0 of the method's type but is already bound by the time you call through `Value.Method`.

## Approach

1. Reject an invalid Value.
2. Resolve the index through the cache; a negative index means `ErrMethod`.
3. `rv.Method(idx).Call` with the single argument and read `out[0].Int()`.

## Solution

```go
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
// 	CallNamed(adder{2}, "Add", 3) => 5, nil
func CallNamed(v any, method string, arg int) (int, error) {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return 0, ErrMethod
	}
	idx := methodIndex(rv.Type(), method)
	if idx < 0 {
		return 0, ErrMethod
	}
	out := rv.Method(idx).Call([]reflect.Value{reflect.ValueOf(arg)})
	return int(out[0].Int()), nil
}
```

## Walkthrough

`adder.Add` has method type `func(adder, int) int`, so `NumIn()` is 2 and `In(1)` is int — it passes. `ptrAdder.Add` is not in `ptrAdder`'s method set at all, so `MethodByName` fails and -1 is cached.

## Pitfalls

- Checking `NumIn() == 1` and rejecting every valid method — the receiver is counted.
- Using `rv.Type().Method(i).Func` and forgetting to pass the receiver as the first argument.
- Caching without the type in the key, so two types with the same method name collide.
