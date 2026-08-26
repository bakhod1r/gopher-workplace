# Deep Copy with Value Receivers

## Intuition

A value receiver gives you a copy of the struct — but **slices** inside the
struct still share the same backing array. Deep cloning means allocating a new
slice and copying elements into it.

## Approach

1. Value receiver already copies `Name` (string is immutable).
2. For `Tags`: if nil, keep nil. Otherwise `make + copy` or `append(nil, ...)`.

## Solution

```go
func (c Config) Clone() Config {
	if c.Tags != nil {
		c.Tags = append([]string(nil), c.Tags...)
	}
	return c
}
```

## Walkthrough

For `Config{"app", ["v1","prod"]}`:
- `c` is already a copy (value receiver).
- `append(nil, ["v1","prod"]...)` → new slice `["v1","prod"]`.
- `c.Tags` now points to the new slice.
- Return `c`.

## Pitfalls

- Just `return c` without cloning `Tags` — shares the backing array.
- Using `make([]string, len(c.Tags))` + `copy` works but is more verbose.
- `append([]string{}, ...)` on nil returns `[]string{}` — test expects nil.
