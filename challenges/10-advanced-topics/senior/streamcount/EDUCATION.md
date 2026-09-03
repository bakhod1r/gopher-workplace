# Count The Lines Without Holding The File

## Intuition

Counting needs one byte of context at a time, so there is no reason to hold the stream. A fixed buffer turns the memory cost from O(input) into O(1).

## Approach

1. Allocate one buffer of a fixed size, e.g. 32 KiB.
2. Loop: `Read` into it, count newlines in `buf[:c]`.
3. Stop on `io.EOF` with a nil error; return other errors with the count so far.

## Solution

```go
import (
	"bytes"
	"io"
)

// CountLines returns the number of '\n' bytes in r.
//
// The reader may deliver gigabytes. The function must work in one pass over
// a fixed-size buffer and must never hold the whole stream in memory.
//
// Examples:
//
// 	CountLines(strings.NewReader("a\nb\n")) => 2, nil
func CountLines(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	n := 0
	for {
		c, err := r.Read(buf)
		n += bytes.Count(buf[:c], []byte{'\n'})
		if err == io.EOF {
			return n, nil
		}
		if err != nil {
			return n, err
		}
	}
}
```

## Walkthrough

A 64 MiB stream is consumed in 2048 reads through one 32 KiB buffer. Total allocation is 32 KiB; `io.ReadAll` would have allocated well over 100 MiB across its doublings.

## Pitfalls

- Ignoring the bytes returned alongside a non-nil error — those are real data.
- Allocating the buffer inside the loop, which turns O(1) memory into O(input) garbage.
