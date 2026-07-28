# Circular indexing

## The idea

A ring buffer stores a logical sequence in a fixed physical array. Logical index
`i` lives at physical `(head + i) mod len`:

```go
return buf[(head+i)%len(buf)]
```

The modulo is what makes the buffer circular; without it, indices run off the end.

## Why it matters

Ring buffers back queues, streaming windows, and fixed-memory logs. The modular
map is the whole abstraction; dropping it turns wrap-around into an out-of-range
panic.

## Watch out

- `%` keeps the index in `[0, len)` for non-negative operands.
- Negative logical indices need the normalize form `((x%n)+n)%n`.
- Overwriting policy (drop-oldest) is separate from the index math.
