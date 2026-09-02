# Ring Buffer Sink

## Intuition

A ring turns "keep the last N" into a fixed allocation made once. Tracking a total write count instead of head/tail pointers makes both the length and the oldest index fall out of arithmetic.

## Approach

1. `Write` stores at `count % len(buf)` and increments `count`.
2. `Len` is `min(count, size)`.
3. `Snapshot` starts at `count - size` when wrapped, else 0, and walks forward modulo the size.

## Solution

```go
func (r *RingSink) Write(line string) {
	r.buf[r.count%len(r.buf)] = line
	r.count++
}

func (r *RingSink) Len() int {
	if r.count < len(r.buf) {
		return r.count
	}
	return len(r.buf)
}

func (r *RingSink) Snapshot() []string {
	n := r.Len()
	out := make([]string, 0, n)

	start := 0
	if r.count > len(r.buf) {
		start = r.count - len(r.buf)
	}
	for i := start; i < r.count; i++ {
		out = append(out, r.buf[i%len(r.buf)])
	}
	return out
}
```

## Walkthrough

After a, b, c, d into a size-3 ring, `count` is 4 and `d` overwrote slot 0. Snapshot starts at index 1 (`4-3`) and walks 1, 2, 3 modulo 3 — slots 1, 2, 0 — giving `[b c d]`.

## Pitfalls

- `append` into the ring, which grows it and defeats the bound.
- Returning `r.buf` directly from `Snapshot`, which is in physical, not logical, order.
- Resetting `count` on wrap, which loses the information needed to find the oldest entry.
