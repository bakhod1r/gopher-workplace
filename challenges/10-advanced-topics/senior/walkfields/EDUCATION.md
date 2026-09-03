# Total Every Int, However Deep

## Intuition

A value tree is walked the same way any tree is: handle the leaf kind, recurse on the container kinds, and stop everywhere else. Reflection turns "what shape is this node" into a `Kind` switch.

## Approach

1. Return 0 for an invalid Value.
2. Int is the leaf: return it.
3. Pointer and interface: return 0 when nil, otherwise recurse on `Elem`.
4. Struct: recurse on every exported field. Slice and array: recurse on every element.

## Solution

```go
import "reflect"

func deepSum(rv reflect.Value) int64 {
	if !rv.IsValid() {
		return 0
	}
	switch rv.Kind() {
	case reflect.Int:
		return rv.Int()
	case reflect.Pointer, reflect.Interface:
		if rv.IsNil() {
			return 0
		}
		return deepSum(rv.Elem())
	case reflect.Struct:
		rt := rv.Type()
		var total int64
		for i := 0; i < rv.NumField(); i++ {
			if !rt.Field(i).IsExported() {
				continue
			}
			total += deepSum(rv.Field(i))
		}
		return total
	case reflect.Slice, reflect.Array:
		var total int64
		for i := 0; i < rv.Len(); i++ {
			total += deepSum(rv.Index(i))
		}
		return total
	default:
		return 0
	}
}

// DeepSum totals every exported int field in v, descending into nested
// structs, slices of structs and pointers.
//
// A nil pointer contributes nothing. Cycles are not part of the input.
//
// Examples:
//
// 	DeepSum(outer{N: 1, In: inner{M: 2}}) => 3
func DeepSum(v any) int64 {
	return deepSum(reflect.ValueOf(v))
}
```

## Walkthrough

For the nested `outer`: N gives 1, In.M gives 2, *Ptr gives 4, and the two list entries give 8 and 16 — 31 in total. `Label` falls through the switch to 0.

## Pitfalls

- Recursing into unexported fields, which panics as soon as something tries to read them out.
- Calling `Elem` on a nil pointer and then `Kind` on the invalid result.
