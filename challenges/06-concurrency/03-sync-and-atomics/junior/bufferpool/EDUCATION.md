# Log Encoder Buffer Pool

## Intuition

A pool is a bucket of scratch space. You borrow a buffer, scribble on it, copy out what you need, and drop it back for the next goroutine. Nothing is shared while you hold it, so no lock is needed.

## Approach

1. Build the pool with `New: func() any { return make([]byte, 0, 256) }`.
2. `Get` a buffer and reset its length with `buf[:0]`.
3. Append each field, with a `|` between neighbours.
4. Convert to `string` (this copies), then `Put` the buffer back.

## Solution

```go
func NewEncoder() *Encoder {
	return &Encoder{
		pool: sync.Pool{
			New: func() any { return make([]byte, 0, 256) },
		},
	}
}

func (e *Encoder) Encode(fields []string) string {
	buf := e.pool.Get().([]byte)
	buf = buf[:0]
	for i, f := range fields {
		if i > 0 {
			buf = append(buf, '|')
		}
		buf = append(buf, f...)
	}
	out := string(buf)
	e.pool.Put(buf) //nolint:staticcheck // slice header copy is fine here
	return out
}
```

## Walkthrough

For `["warn", "disk full"]`: the pool hands out an empty slice; `warn` is appended, then `|`, then `disk full`; `string(buf)` copies the bytes into an immutable string; the buffer goes back to the pool with its capacity intact.

## Pitfalls

- Returning `string` **before** `Put` is required — after `Put`, another goroutine may overwrite the bytes.
- Forgetting `buf[:0]` leaks the previous line into the new one.
- Assuming an object survives in the pool: `New` must be able to produce one at any time.
