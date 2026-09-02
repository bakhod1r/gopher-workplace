# Get Or Default

## Intuition

Reading a missing key returns the zero value silently. The second return value is the presence flag, and it is the whole puzzle here.

## Approach

1. Look up `k` with the comma-ok form.
2. Return the value when `ok` is true.
3. Return `def` otherwise.

## Solution

```go
func GetOr[K comparable, V any](m map[K]V, k K, def V) V {
	if v, ok := m[k]; ok {
		return v
	}
	return def
}
```

## Walkthrough

`GetOr(map[string]int{"a": 0}, "a", 9)` finds `ok == true`, so it returns the stored `0`, not the default `9`.

## Pitfalls

- Comparing `m[k]` to the zero value instead of using `ok` — that overrides legitimately stored zeros.
- Returning `def` when the map itself is nil but the key logic would still work — a nil map lookup is legal and yields `ok == false`.
- Declaring `V comparable` unnecessarily; nothing here compares values.
