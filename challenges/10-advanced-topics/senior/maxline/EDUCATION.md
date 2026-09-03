# Longest Line, Bounded Working Set

## Intuition

The temptation is to accumulate the line so you can measure it. But the length is a running counter — it survives buffer boundaries for free, and the line itself never needs to be stored.

## Approach

1. Keep `cur` (current line length) and `best` across reads.
2. For each byte: on `\n`, fold `cur` into `best` and reset; otherwise increment `cur`.
3. At EOF fold the final `cur` and return `best`.

## Solution

```go
import "io"

// MaxLine returns the length in bytes of the longest '\n'-separated line
// in r, not counting the newline itself.
//
// Lines may be longer than any single read, and the stream may be far
// larger than memory. Only a fixed-size buffer may be held.
//
// Examples:
//
// 	MaxLine(strings.NewReader("ab\ncdef\n")) => 4, nil
func MaxLine(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	best, cur := 0, 0
	for {
		c, err := r.Read(buf)
		for _, b := range buf[:c] {
			if b == '\n' {
				if cur > best {
					best = cur
				}
				cur = 0
				continue
			}
			cur++
		}
		if err == io.EOF {
			if cur > best {
				best = cur
			}
			return best, nil
		}
		if err != nil {
			return 0, err
		}
	}
}
```

## Walkthrough

A 32 MiB single line is consumed in 1024 reads through one 32 KiB buffer; `cur` reaches 33554432 and is folded into `best` at EOF. Nothing beyond the buffer is ever allocated.

## Pitfalls

- Resetting `cur` at the start of each read instead of at each newline.
- Forgetting the unterminated final line.
