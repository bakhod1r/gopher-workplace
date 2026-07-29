# Capturing shared loop variables

## Intuition

A variable declared outside the loop is shared by every closure; only per-iteration variables (or an explicit inner copy) give each closure its own value.

## Approach

1. A single hoisted `var i` is shared by every closure, so all capture the final `i`.
2. Use `for i := range names` so each iteration has its own `i` (Go 1.22 per-iteration scope).

## Solution

```go
func Labelers(names []string) []func() string {
	var out []func() string
	for i := range names {
		out = append(out, func() string { return names[i] })
	}
	return out
}
```

## Walkthrough

With the shared `i`, every closure reads the loop's end value and indexes the same element. A per-iteration `i` binds each closure to its own index.

## Pitfalls

- `var i int; for i = 0; ...` shares one `i` across all closures.
- Use `for i := range` or copy into a fresh variable inside the loop.
