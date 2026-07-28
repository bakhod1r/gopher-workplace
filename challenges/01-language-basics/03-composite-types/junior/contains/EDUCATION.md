# Membership by linear scan

## The idea

For an unsorted slice, membership is a linear scan with an early return:

```go
for _, x := range xs {
	if x == target { return true }
}
return false
```

## Why it matters

It is the simplest search and the basis for filtering and dedup. For large or
hot lookups you'd switch to a `map[T]struct{}` set (O(1)); for small slices a
scan is fine and allocation-free.

## Why the standard library

`slices.Contains` (Go 1.21+) does exactly this; hand-writing it shows the
early-return pattern.

## Watch out

- Return `false` only after the whole loop — not inside it.
- `==` must be defined for the element type (strings, numbers, comparable
  structs).
- Ranging nil is safe and yields `false`.
