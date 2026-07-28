# Editing strings

## The idea

A Go string is immutable — you cannot assign to `s[i]`. To transform it, copy to a
`[]byte`, edit that, and convert back:

```go
b := []byte(s)
// mutate b...
return string(b)
```

Returning the original `s` discards all the work.

## Why it matters

String immutability is a core guarantee (safe sharing, map keys, no defensive
copies). The transform pattern (`[]byte` round-trip) is how you "change" a
string — and forgetting the final `string(b)` is a silent no-op.

## Watch out

- `s[i] = x` is a compile error; strings are read-only.
- Each conversion (`[]byte(s)`, `string(b)`) copies.
- For rune-level edits, use `[]rune`; for ASCII, `[]byte` suffices.
