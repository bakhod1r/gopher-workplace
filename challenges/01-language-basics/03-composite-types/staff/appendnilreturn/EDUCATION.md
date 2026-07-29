# The nil slice is usable

## Intuition

A nil slice has length 0 and is safe to `range`, `len`, and — crucially —
`append`:

```go
var s []int      // nil
s = append(s, 5) // [5], allocated on demand
```

So a `if s == nil { return nil }` guard before an append is not just unnecessary,
it's a bug: it refuses to add the first element.

## Approach

1. Bug: the guard `if s == nil { return nil }` short-circuits and returns nil for a nil input, dropping x entirely.
2. Fix: delete the guard. `append(s, x)` already handles a nil s correctly by allocating a fresh backing array.

## Solution

```go
func Add(s []int, x int) []int {
	return append(s, x)
}
```

## Walkthrough

Call Add(nil, 5): with the bug, s==nil is true so it returns nil (wrong). After the fix the guard is gone, so append(nil, 5) allocates [5] and returns it.

## Pitfalls

- Append to nil: fine. Write to nil map: panic.
- `len(nil) == 0`, ranging nil is zero iterations.
- Distinguish nil from empty only when serialization/identity demands it.
