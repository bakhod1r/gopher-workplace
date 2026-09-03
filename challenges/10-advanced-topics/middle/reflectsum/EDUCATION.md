# Total Any Slice Of Integers

## Intuition

Reflection collapses "one branch per integer width" into a single loop, because `Int()` normalises every signed kind to int64. The type switch a caller would have written becomes two kind checks.

## Approach

1. Reject anything that is not a slice or array.
2. Reject an element kind that is not a signed integer.
3. Loop with `rv.Index(i).Int()` into an int64.

## Solution

```go
import "reflect"

// Sum totals v when it is a slice or array of a signed integer kind, and
// reports whether it could.
//
// Examples:
//
// 	Sum([]int32{1, 2}) => 3, true
func Sum(v any) (int64, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return 0, false
	}
	switch rv.Type().Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	default:
		return 0, false
	}
	var total int64
	for i := 0; i < rv.Len(); i++ {
		total += rv.Index(i).Int()
	}
	return total, true
}
```

## Walkthrough

`[]myInt{2,3}` has element kind int, so `Int()` reads each element and the total is 5 — with no mention of `myInt` anywhere.

## Pitfalls

- Calling `Int()` on an unsigned element, which panics — `Uint()` is the accessor for those.
- Accumulating in the element's own width, which overflows for large int32 values.
