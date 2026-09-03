# The Scanner That Gave Up On A Long Line

## Intuition

`bufio.Scanner` caps token size on purpose, so a malformed stream cannot exhaust memory. The default is a guess about your data, and when it is wrong the loop simply stops and reports the error afterwards.

## Approach

1. Call `sc.Buffer` with an initial buffer and `maxLine` as the cap.
2. Scan as before and check `sc.Err()`.

## Solution

```go
import (
	"bufio"
	"io"
)

// maxLine is the longest line this reader must accept.
const maxLine = 4 << 20

// LongestLine returns the length of the longest line in r.
//
// bufio.Scanner refuses tokens larger than its buffer limit, which defaults
// to 64 KiB. A line longer than that is an error, not a truncation.
//
// Examples:
//
// 	LongestLine(strings.NewReader("ab\ncdef")) => 4, nil
func LongestLine(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), maxLine)
	best := 0
	for sc.Scan() {
		if n := len(sc.Bytes()); n > best {
			best = n
		}
	}
	if err := sc.Err(); err != nil {
		return 0, err
	}
	return best, nil
}
```

## Walkthrough

A 200 KiB line exceeds the default 64 KiB, so `Scan` returns false immediately and `Err` is `bufio.ErrTooLong`. Raising the cap to 4 MiB lets the same input through.

## Pitfalls

- Ignoring `sc.Err()`, which turns the failure into a silently short answer.
- Removing the cap entirely by passing a huge maximum — the limit is a safety feature.
