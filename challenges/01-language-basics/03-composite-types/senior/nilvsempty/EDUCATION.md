# Nil and empty slices differ at the boundary

## The idea

`var out []string` is **nil**; `[]string{}` is non-nil with length 0. They behave
identically for `len`, `range`, and `append` — but not for identity or
serialization:

```go
out := []string{} // marshals to [], not null
```

## Why it matters

APIs and clients often distinguish `[]` (present, empty) from `null` (absent). A
handler that returns a nil slice emits `null`, breaking consumers that expect an
array. The contract, not the length, dictates which to use.

## Watch out

- `out == nil` is the only way to tell them apart in Go.
- `append` promotes nil to non-nil, so the distinction only matters when nothing
  is appended.
- Prefer `[]T{}` when the emptiness must be observable.
