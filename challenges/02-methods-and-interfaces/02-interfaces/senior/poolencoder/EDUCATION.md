# Pooled Encoder

## Intuition

`sync.Pool` recycles memory the GC would otherwise churn through. Two rules make it safe: reset before use, and never hand out anything that still points into the pooled buffer.

## Approach

1. `Get` a `*[]byte` from the pool and reslice it to zero length.
2. Append fields and separators into it.
3. Convert with `string(buf)` — this copies, so the result is independent.
4. Store the (possibly regrown) buffer back into the pointer and `Put` it.

## Solution

```go
func (e *PooledEncoder) Encode(fields []string) string {
	bp := e.pool.Get().(*[]byte)
	buf := (*bp)[:0]

	for i, f := range fields {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, f...)
	}
	out := string(buf)

	*bp = buf
	e.pool.Put(bp)
	return out
}
```

## Walkthrough

Pooling `*[]byte` rather than `[]byte` matters: `append` may regrow the slice, and only the pointer indirection lets the larger buffer make it back into the pool.

## Pitfalls

- Returning `unsafe`-style aliases of `buf` — the next `Encode` overwrites the caller's data.
- Forgetting `[:0]`, so each call appends onto the previous contents.
- Pooling `[]byte` by value, which silently discards any growth and defeats the reuse.
