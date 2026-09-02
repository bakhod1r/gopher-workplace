# Translation Queue

## Intuition

Counting one message is entirely local work. The only shared memory is the
output slice, and each goroutine touches exactly one distinct slot of it.

## Approach

1. Allocate `out := make([]int, len(messages))`.
2. Launch one goroutine per message, passing `i` and the message.
3. Write `utf8.RuneCountInString(msg)` to `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package translationqueue — Gopher Workplace challenge.
package translationqueue

import (
	"sync"
	"unicode/utf8"
)

// CharCounts returns the billable character count of every queued message.
//
// Examples:
//
//	CharCounts([]string{"go", "añb"})  => [2 3]
//	CharCounts([]string{"日本"})        => [2]
//	CharCounts(nil)                    => []
func CharCounts(messages []string) []int {
	out := make([]int, len(messages))
	var wg sync.WaitGroup
	for i, msg := range messages {
		wg.Add(1)
		go func(i int, msg string) {
			defer wg.Done()
			out[i] = utf8.RuneCountInString(msg)
		}(i, msg)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"añb"` is four bytes but three characters, so the count is `3`.
- `"日本"` is six bytes and two characters.
- `"a🙂"` is five bytes and two characters — billing would be wrong with `len`.

## Pitfalls

- Using `len(msg)`, which overcharges every message containing non-ASCII text.
- Allocating `[]rune(msg)` only to take its length — correct but wasteful.
- Forgetting the `unicode/utf8` import.
