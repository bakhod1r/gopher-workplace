# Windowing a slice

## The idea

Chunking steps an index by `size` and takes a sub-slice each time, clamping the
end so the final (short) chunk fits:

```go
for i := 0; i < len(xs); i += size {
	end := i + size
	if end > len(xs) { end = len(xs) }
	out = append(out, xs[i:end])
}
```

## Why it matters

Batching (DB inserts, API pages, worker tasks) needs fixed-size groups with a
graceful remainder. Clamping the end is what handles the last partial chunk.

## Watch out

- The chunks **share** `xs`'s backing array (sub-slices, not copies).
- Guard `size <= 0` to avoid an infinite loop.
- `min(i+size, len(xs))` (Go 1.21+) replaces the manual clamp.
