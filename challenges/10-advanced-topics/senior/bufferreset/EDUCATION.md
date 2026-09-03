# The Shared Buffer Nobody Emptied

## Intuition

Sharing a buffer to save allocations means inheriting its contents. The lock stops two callers writing at once; it does nothing about the bytes the previous caller left behind.

## Approach

1. Take the lock as before.
2. `scratch.Reset()` before writing.
3. Render and return the string.

## Solution

```go
import (
	"bytes"
	"strconv"
	"sync"
)

var (
	mu      sync.Mutex
	scratch bytes.Buffer
)

// Render returns vals as decimal numbers joined by '-'.
//
// The package keeps one scratch buffer to avoid allocating per call. A
// shared buffer has to be emptied before it is written to.
//
// Examples:
//
// 	Render([]int{1, 2}) => "1-2"
func Render(vals []int) string {
	mu.Lock()
	defer mu.Unlock()
	scratch.Reset()
	for i, v := range vals {
		if i > 0 {
			scratch.WriteByte('-')
		}
		scratch.WriteString(strconv.Itoa(v))
	}
	return scratch.String()
}
```

## Walkthrough

The first call writes "7" and leaves it there. The second appends, returning "77". After 2000 calls the buffer holds thousands of bytes and every response carries all of them.

## Pitfalls

- Resetting after `String()` instead of before writing — a concurrent caller may already be inside the lock's queue with stale expectations.
- Removing the shared buffer entirely, which fixes the symptom and gives up the optimisation.
