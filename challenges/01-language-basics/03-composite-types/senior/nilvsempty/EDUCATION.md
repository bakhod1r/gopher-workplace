# Nil and empty slices differ at the boundary

## Intuition

`var out []string` is **nil**; `[]string{}` is non-nil with length 0. They behave
identically for `len`, `range`, and `append` — but not for identity or
serialization:

```go
out := []string{} // marshals to [], not null
```

## Approach

1. Bug: var out []string is nil; when nothing matches it stays nil and JSON-encodes as null instead of []. 2. Fix: initialize out := []string{} — a non-nil empty slice. 3. Appends work identically; the difference is the empty case returns [] not null.

## Solution

```go
func NonEmpty(in []string) []string {
	out := []string{}
	for _, s := range in {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
```

## Walkthrough

Input with no non-empty strings: buggy out is nil (encodes null). With []string{} out is empty-but-non-nil (encodes []).

## Pitfalls

- `out == nil` is the only way to tell them apart in Go.
- `append` promotes nil to non-nil, so the distinction only matters when nothing
  is appended.
- Prefer `[]T{}` when the emptiness must be observable.
