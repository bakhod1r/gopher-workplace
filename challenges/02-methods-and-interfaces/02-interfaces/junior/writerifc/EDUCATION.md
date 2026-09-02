# Writer Interface

## Intuition

A writer that reports its byte count lets callers total the output without inspecting the sink — the same idea behind `io.Writer`.

## Approach

1. `Write` appends to `b.buf` and returns `len(s)`.
2. `String` returns `b.buf`.
3. `WriteLines` writes `line + "\n"` and adds each returned count to a total.

## Solution

```go
func (b *Builder) Write(s string) int {
	b.buf += s
	return len(s)
}

func (b *Builder) String() string { return b.buf }

func WriteLines(w Writer, lines []string) int {
	total := 0
	for _, line := range lines {
		total += w.Write(line + "\n")
	}
	return total
}
```

## Walkthrough

`["a", "b"]` writes `"a\n"` (2 bytes) and `"b\n"` (2 bytes) — total 4, buffer `"a\nb\n"`.

## Pitfalls

- Returning `len(b.buf)` instead of `len(s)` — that reports the running size, not this write.
- Forgetting the newline, which makes the total 2 instead of 4.
- A value receiver on `Write`, so nothing accumulates.
