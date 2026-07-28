# Writing into a buffer at an offset

## The idea

Filling a preallocated buffer means copying each piece to its correct offset. The
second slice starts where the first ends:

```go
out := make([]byte, len(a)+len(b))
copy(out, a)
copy(out[len(a):], b) // offset by len(a)
```

## Why it matters

Manual buffer assembly (encoders, network frames, `bytes` plumbing) depends on
offset arithmetic. Copying to offset 0 twice overwrites earlier data — a classic
buffer-building bug.

## Watch out

- `copy` returns the count; the destination sub-slice sets where it lands.
- Size the buffer exactly, or trailing zeros/overflows appear.
- For many pieces, `bytes.Buffer` or `append` is simpler than manual offsets.
