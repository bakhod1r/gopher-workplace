# Map Lookup

## Intuition

A map read always succeeds — it hands back the zero value for a missing key. Only the comma-ok form tells you whether the key was really there, and that distinction is what the error reports.

## Approach

1. Read with `v, ok := m[key]`.
2. Return `0, ErrNotFound` when `ok` is false.
3. Otherwise return `v, nil`.

## Solution

```go
v, ok := m[key]
if !ok {
	return 0, ErrNotFound
}
return v, nil
```

## Walkthrough

For key `"zero"` the stored value is 0 but `ok` is true, so the result is `0, nil` — not an error.

## Pitfalls

- Testing `if m[key] == 0` — that misreports a legitimately stored zero.
- Guarding against a nil map by hand; nil map reads are already safe.
- Returning the value alongside the error when the key is missing.
