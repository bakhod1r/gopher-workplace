# Draining A Reader Into The Caller's Buffer

## Intuition

Keep the data in one slice whose length tracks how much has been read. When there is no spare capacity, grow once and carry on.

## Approach

1. Reset `dst` and normalise `chunk`.
2. Loop: ensure `chunk` bytes of spare capacity, read into the spare region, extend the length.
3. Stop on EOF; return any other error along with the data.

## Solution

```go
func ReadAll(r io.Reader, dst []byte, chunk int) ([]byte, error) {
	if chunk <= 0 {
		chunk = 4096
	}
	dst = dst[:0]
	for {
		if cap(dst)-len(dst) < chunk {
			grown := make([]byte, len(dst), max(2*cap(dst), len(dst)+chunk))
			copy(grown, dst)
			dst = grown
		}
		n, err := r.Read(dst[len(dst) : len(dst)+chunk])
		dst = dst[:len(dst)+n]
		if errors.Is(err, io.EOF) {
			return dst, nil
		}
		if err != nil {
			return dst, err
		}
	}
}

func CountChunks(n, chunk int) int {
	if n <= 0 || chunk <= 0 {
		return 0
	}
	return (n+chunk-1)/chunk + 1
}
```

## Walkthrough

`dst[len(dst) : len(dst)+chunk]` is a window over the spare capacity — the read writes there, and extending the length publishes exactly the bytes that arrived.

## Pitfalls

- Treating `io.EOF` as an error, so every successful read reports failure.
- Ignoring `n` when `err != nil`; a reader may return data *and* an error in the same call.
- Discarding the partial data on error, which throws away everything already transferred.
