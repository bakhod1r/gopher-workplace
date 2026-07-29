# Pointer values in maps

## Intuition

Storing `*int` as map values lets updates flow to the underlying variables, unlike storing ints by value.

## Approach

1. Range over the map's values (the pointers).
2. `*p++` mutates each pointee.
3. The map itself is unchanged — only the pointed-at ints move.

## Solution

```go
func BumpAll(m map[string]*int) {
	for _, p := range m {
		*p++
	}
}
```

## Walkthrough

With `{"a": &a}` and `a = 1`: the loop dereferences `&a` and increments, so `a` becomes `2`.

## Pitfalls

- Map iteration order is random but irrelevant here.
- `*p++` changes the referenced variable.
