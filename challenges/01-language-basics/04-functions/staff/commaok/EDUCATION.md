# The comma-ok idiom for maps

## Intuition

`v, ok := m[k]` is the only way to tell a stored zero from a missing key; the single-value form collapses them.

## Approach

1. A plain map read can't distinguish a missing key from a zero value.
2. Use the comma-ok form: `score, ok = scores[name]`.

## Solution

```go
func Lookup(scores map[string]int, name string) (score int, ok bool) {
	score, ok = scores[name]
	return
}
```

## Walkthrough

The bug always sets `ok = true`, so a missing key looks present. Comma-ok reports `false` for absent keys while still returning `0, true` for a real zero.

## Pitfalls

- `m[k]` alone can't distinguish missing from zero-valued.
- Always use `v, ok := m[k]` when absence matters.
