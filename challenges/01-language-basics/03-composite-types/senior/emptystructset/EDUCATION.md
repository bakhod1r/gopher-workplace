# Bool sets store true

## The idea

When a `map[T]bool` is used as a set and membership is tested by the value
(`if inB[x]`), members must be stored as **true**:

```go
inB[x] = true
```

Storing `false` makes every lookup report "absent".

## Why it matters

`map[T]bool` sets are convenient but couple the marker value to the test. The
`struct{}` + comma-ok style (`_, ok := set[x]`) avoids the trap entirely by
testing presence, not value.

## Watch out

- With value-based tests, the stored bool must be `true`.
- `map[T]struct{}` uses no value memory and tests via `ok`.
- Reading a missing key gives `false` — indistinguishable from a stored `false`.
