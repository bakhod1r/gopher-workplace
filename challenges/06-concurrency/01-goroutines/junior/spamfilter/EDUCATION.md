# Spam Filter

## Intuition

'Filter' is awkward to parallelise safely, but 'mark' is trivial. Return a
boolean per message and the race disappears along with the ordering problem.

## Approach

1. Allocate `out := make([]bool, len(messages))`.
2. Launch one goroutine per message, passing `i` and the message.
3. Write `strings.Contains(msg, banned)` to `out[i]`.
4. `wg.Wait()` before returning.

## Solution

```go
// Package spamfilter — Gopher Workplace challenge.
package spamfilter

import (
	"strings"
	"sync"
)

// Flagged marks every message that contains the banned phrase.
//
// Examples:
//
//	Flagged([]string{"buy now", "hello"}, "buy")  => [true false]
//	Flagged([]string{"hello"}, "buy")             => [false]
//	Flagged(nil, "buy")                           => []
func Flagged(messages []string, banned string) []bool {
	out := make([]bool, len(messages))
	var wg sync.WaitGroup
	for i, msg := range messages {
		wg.Add(1)
		go func(i int, msg string) {
			defer wg.Done()
			out[i] = strings.Contains(msg, banned)
		}(i, msg)
	}
	wg.Wait()
	return out
}
```

## Walkthrough

- `"buy now"` contains `"buy"`, so its verdict is `true`.
- `"rebuy"` contains it as a substring too, so it is flagged as well.
- An empty message contains nothing and is not flagged.

## Pitfalls

- `append`-ing flagged messages from goroutines — a race with a nondeterministic order.
- Expecting `strings.Contains(s, "")` to be false; it is always true.
- Capturing `msg` from the loop variable instead of passing it as a parameter.
