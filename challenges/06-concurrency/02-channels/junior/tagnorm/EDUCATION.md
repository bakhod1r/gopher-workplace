# A Pipeline Stage

## Intuition

Every pipeline stage has the same skeleton: read until the upstream
closes, write transformed values downstream, then close downstream. Close
propagates termination along the chain like a domino.

## Approach

1. `range` over `in`.
2. Send `strings.ToUpper(tag)` on `out` for each value.
3. `close(out)` after the loop.

## Solution

```go
import "strings"

func NormalizeTags(in <-chan string, out chan<- string) {
	for tag := range in {
		out <- strings.ToUpper(tag)
	}
	close(out)
}
```

## Walkthrough

For `"az"`, `"eu"`: send `"AZ"`, send `"EU"`, the receiver closes so the
loop exits, then `close(out)` tells the writer the stage is done.

## Pitfalls

- Closing `in` — a stage must never close its input; the compiler blocks it for `<-chan` anyway.
- Forgetting `close(out)` leaves the writer's `range` hanging.
- If `out` is unbuffered and nobody receives, this stage must run in its own goroutine.
