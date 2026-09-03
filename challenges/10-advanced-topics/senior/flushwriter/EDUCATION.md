# The Buffered Writer Nobody Flushed

## Intuition

A buffered writer trades immediate writes for batched ones, which means the last batch only leaves when you say so. It is also the point at which the underlying writer's failures finally become visible.

## Approach

1. Write the lines as before.
2. Return `bw.Flush()` instead of nil.

## Solution

```go
import (
	"bufio"
	"io"
)

// WriteAll writes each line followed by '\n' through a buffered writer.
//
// A buffered writer holds the tail of the output until it is flushed; the
// last partial buffer is lost otherwise.
//
// Examples:
//
// 	WriteAll(&buf, []string{"a"}) => buf holds "a\n"
func WriteAll(w io.Writer, lines []string) error {
	bw := bufio.NewWriter(w)
	for _, l := range lines {
		if _, err := bw.WriteString(l); err != nil {
			return err
		}
		if err := bw.WriteByte('\n'); err != nil {
			return err
		}
	}
	return bw.Flush()
}
```

## Walkthrough

500 lines of 41 bytes is 20500, and the default buffer is 4096 — so 20480 bytes leave on their own and the last 20 sit in the buffer forever without the flush.

## Pitfalls

- `defer bw.Flush()` — it flushes, and it discards the error.
- Assuming a small output is safe; a small output is entirely inside the buffer.
