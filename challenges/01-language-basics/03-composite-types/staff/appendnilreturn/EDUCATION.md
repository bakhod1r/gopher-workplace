# The nil slice is usable

## The idea

A nil slice has length 0 and is safe to `range`, `len`, and — crucially —
`append`:

```go
var s []int      // nil
s = append(s, 5) // [5], allocated on demand
```

So a `if s == nil { return nil }` guard before an append is not just unnecessary,
it's a bug: it refuses to add the first element.

## Why it matters

The nil slice as "empty" is a deliberate Go convenience. Over-defensive nil checks
around append add bugs and noise. (Nil *maps* differ: writing to a nil map
panics.)

## Watch out

- Append to nil: fine. Write to nil map: panic.
- `len(nil) == 0`, ranging nil is zero iterations.
- Distinguish nil from empty only when serialization/identity demands it.
