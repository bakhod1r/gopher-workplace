# Map lookup

## Intuition

Indexing a map returns the value, or the value type's zero if the key is absent.
The comma-ok form reports presence:

```go
v, ok := m[key]
```

## Approach

1. Read v, ok := m[key] — the comma-ok map access.
2. Return v, ok directly. On a nil map the read yields 0, false without panicking.

## Solution

```go
func Lookup(m map[string]int, key string) (int, bool) {
	v, ok := m[key]
	return v, ok
}
```

## Walkthrough

Lookup({"z":0},"z"): comma-ok returns v=0, ok=true, so the caller learns the key exists despite the zero value.

## Pitfalls

- Reading a nil map is safe (returns zero); writing panics.
- Iteration order is randomized.
- Keys must be comparable.
