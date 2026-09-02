# Chunked Reader

## Intuition

The `io.Reader` shape exists so the caller controls memory. The reader fills whatever buffer it is handed, so a 100M-byte file costs one buffer, not 100M bytes.

## Approach

1. `Read` returns `(0, false)` when drained or given an empty buffer.
2. Otherwise `copy(p, s.Data[s.pos:])` — `copy` stops at the shorter of the two.
3. Advance `pos` by the copied count and return it.
4. `CountLines` loops, scanning only `buf[:n]` for `'\n'`.

## Solution

```go
func (s *ChunkSource) Read(p []byte) (int, bool) {
	if s.pos >= len(s.Data) || len(p) == 0 {
		return 0, false
	}
	n := copy(p, s.Data[s.pos:])
	s.pos += n
	return n, true
}

func CountLines(r Reader, buf []byte) int {
	lines := 0
	for {
		n, ok := r.Read(buf)
		if !ok {
			return lines
		}
		for _, b := range buf[:n] {
			if b == '\n' {
				lines++
			}
		}
	}
}
```

## Walkthrough

With a 2-byte buffer over `"aaa\nbbb\n"`, the newline at index 3 lands in the second chunk — counting per chunk still totals 2 because the counter lives outside the loop.

## Pitfalls

- Scanning all of `buf` instead of `buf[:n]`, which counts stale bytes from the previous chunk.
- Allocating a fresh buffer inside the loop, which is exactly the cost this design removes.
- Returning `len(p)` from `Read` instead of the `copy` result, overreporting on the final partial chunk.
