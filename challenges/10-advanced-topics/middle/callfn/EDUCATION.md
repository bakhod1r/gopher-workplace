# Call A Function You Only Know At Run Time

## Intuition

`Call` is unforgiving: the wrong arity or the wrong argument type is a panic, not an error. The function's type carries everything needed to check the shape before you call, which turns a crash into a returned error.

## Approach

1. Verify the kind is func, it is not variadic, and the arity matches.
2. Verify every parameter and every result is of kind int.
3. Build the `[]reflect.Value` arguments, `Call`, and read the results with `Int()`.

## Solution

```go
import (
	"errors"
	"reflect"
)

// ErrSignature is returned when fn does not match the expected shape.
var ErrSignature = errors.New("fn must take and return only ints")

// CallInts calls fn with args and returns its int results.
//
// fn must be a function taking exactly len(args) int parameters and
// returning only ints. Anything else is an error, not a panic.
//
// Examples:
//
// 	CallInts(func(a, b int) int { return a + b }, 1, 2) => []int{3}
func CallInts(fn any, args ...int) ([]int, error) {
	rv := reflect.ValueOf(fn)
	if rv.Kind() != reflect.Func {
		return nil, ErrSignature
	}
	t := rv.Type()
	if t.IsVariadic() || t.NumIn() != len(args) {
		return nil, ErrSignature
	}
	for i := 0; i < t.NumIn(); i++ {
		if t.In(i).Kind() != reflect.Int {
			return nil, ErrSignature
		}
	}
	for i := 0; i < t.NumOut(); i++ {
		if t.Out(i).Kind() != reflect.Int {
			return nil, ErrSignature
		}
	}
	in := make([]reflect.Value, len(args))
	for i, a := range args {
		in[i] = reflect.ValueOf(a)
	}
	res := rv.Call(in)
	out := make([]int, len(res))
	for i, r := range res {
		out[i] = int(r.Int())
	}
	return out, nil
}
```

## Walkthrough

For `func(a, b int) int` with args 1 and 2: arity 2 matches, both parameters are int, the single result is int. `Call` returns one Value holding 3.

## Pitfalls

- Checking arity but not parameter types — `Call` panics on a type mismatch.
- Forgetting the variadic case, where `NumIn` counts the slice parameter as one.
