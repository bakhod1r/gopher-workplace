# Clone A Map Without Reflection

## Intuition

Generics give one implementation the compiler specialises per instantiation — the same code a hand-written cloner would be, without the four copies or the boxing a reflective version needs.

## Approach

1. Return nil for a nil map.
2. `make(map[K]V, len(m))`.
3. Copy every entry and return.

## Solution

```go
// CloneMap returns a shallow copy of m with the same entries.
//
// A type parameter keeps the keys and values concrete, so nothing is boxed
// and the copy costs one allocation plus the entries.
//
// Examples:
//
// 	CloneMap(map[string]int{"a": 1}) => a new map with the same entry
func CloneMap[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}
	out := make(map[K]V, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
```

## Walkthrough

Cloning a 64-entry map allocates one sized map; the entries are copied by value, so a struct value is independent and a slice value is shared.

## Pitfalls

- Returning an empty map for a nil input, which changes `== nil` behaviour downstream.
- Omitting the size hint and rehashing on the way in.
- Documenting it as a deep copy; nested slices and maps are shared.
