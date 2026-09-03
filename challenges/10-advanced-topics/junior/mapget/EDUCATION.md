# Missing Or Zero

## Intuition

A map index always produces a value — the zero value when the key is missing. The second result is the only thing that distinguishes "stored zero" from "not there", which is exactly the distinction most map bugs turn on.

## Approach

1. `v, ok := m[key]` and return both.

## Solution

```go
// Get returns the value stored under key and whether the key was present.
//
// A missing key reads as 0, which is also a value a key can hold — only the
// second result tells them apart.
//
// Examples:
//
// 	Get(map[string]int{"a": 0}, "a") => 0, true
func Get(m map[string]int, key string) (int, bool) {
	v, ok := m[key]
	return v, ok
}
```

## Walkthrough

Reading a nil map yields the zero value and false, so no separate nil check is needed.

## Pitfalls

- `if m[key] != 0` as a presence test, which is the bug this closes.
- Writing to a nil map, which does panic — only reads are safe.
