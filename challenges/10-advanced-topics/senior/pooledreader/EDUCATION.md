# The Pooled Reader Still Reading The Last Request

## Intuition

A pool recycles objects with their state intact. A buffered reader's state includes the source it wraps and whatever it has already read ahead — so a borrowed one keeps serving the previous request until it is rebound.

## Approach

1. Get the reader from the pool.
2. `br.Reset(r)` to bind it to this call's source.
3. Read the line, trim the newline, tolerate `io.EOF`.

## Solution

```go
import (
	"bufio"
	"io"
	"strings"
	"sync"
)

var pool = sync.Pool{New: func() any { return bufio.NewReaderSize(nil, 64) }}

// FirstLine returns the first '\n'-terminated line from r, using a
// pooled bufio.Reader.
//
// A bufio.Reader taken from a pool is still attached to the previous
// source, with the previous source's buffered bytes inside it.
//
// Examples:
//
// 	FirstLine(strings.NewReader("a\nb\n")) => "a", nil
func FirstLine(r io.Reader) (string, error) {
	br := pool.Get().(*bufio.Reader)
	br.Reset(r)
	defer pool.Put(br)
	line, err := br.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSuffix(line, "\n"), nil
}
```

## Walkthrough

The first call works because `New` builds a reader over a nil source, which fails and is then... not what happens: the reader is never bound, so every call reads from whatever the last `Reset` left — and with no `Reset` at all, from nil.

## Pitfalls

- Resetting after reading — the read has already happened.
- Putting the reader back still bound to the caller's source, which keeps that source alive; `Reset(nil)` before `Put` avoids it.
