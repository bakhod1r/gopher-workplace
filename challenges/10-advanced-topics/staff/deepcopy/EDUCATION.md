# A Copy That Shares Nothing

## Intuition

A shallow copy duplicates the headers and shares the storage. A deep copy has to rebuild every container it meets, which means one construction rule per container kind — and a decision about nil, which is a distinct state from empty.

## Approach

1. Return nil for an invalid Value.
2. Delegate to the recursive helper.
3. Box the result back with `Interface()`.

## Solution

```go
import "reflect"

func deepCopy(rv reflect.Value) reflect.Value {
	switch rv.Kind() {
	case reflect.Pointer:
		if rv.IsNil() {
			return rv
		}
		out := reflect.New(rv.Type().Elem())
		out.Elem().Set(deepCopy(rv.Elem()))
		return out
	case reflect.Interface:
		if rv.IsNil() {
			return rv
		}
		out := reflect.New(rv.Type()).Elem()
		out.Set(deepCopy(rv.Elem()))
		return out
	case reflect.Slice:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeSlice(rv.Type(), rv.Len(), rv.Len())
		for i := 0; i < rv.Len(); i++ {
			out.Index(i).Set(deepCopy(rv.Index(i)))
		}
		return out
	case reflect.Array:
		out := reflect.New(rv.Type()).Elem()
		for i := 0; i < rv.Len(); i++ {
			out.Index(i).Set(deepCopy(rv.Index(i)))
		}
		return out
	case reflect.Map:
		if rv.IsNil() {
			return rv
		}
		out := reflect.MakeMapWithSize(rv.Type(), rv.Len())
		iter := rv.MapRange()
		for iter.Next() {
			out.SetMapIndex(deepCopy(iter.Key()), deepCopy(iter.Value()))
		}
		return out
	case reflect.Struct:
		out := reflect.New(rv.Type()).Elem()
		for i := 0; i < rv.NumField(); i++ {
			if rv.Type().Field(i).IsExported() {
				out.Field(i).Set(deepCopy(rv.Field(i)))
			}
		}
		return out
	default:
		return rv
	}
}

// DeepCopy returns a copy of v that shares no mutable storage with it.
//
// Structs, slices, maps, arrays and pointers are copied recursively; scalars
// and strings are copied by value. Cycles are not part of the input.
//
// Examples:
//
// 	DeepCopy(node{Tags: []string{"a"}}) => an independent node
func DeepCopy(v any) any {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return nil
	}
	return deepCopy(rv).Interface()
}
```

## Walkthrough

Copying `node{Tags:["a"], Child:&node{...}}` makes a new two-element `[]string`, a new `node` for the child, and a new pointer to it. Writing through any of them cannot reach the original.

## Pitfalls

- Turning a nil slice into an empty one, which breaks `== nil` checks downstream.
- Copying map values without copying keys — keys can be structs containing pointers too.
- Forgetting that unexported fields cannot be set, so a copy of such a struct is necessarily partial.
