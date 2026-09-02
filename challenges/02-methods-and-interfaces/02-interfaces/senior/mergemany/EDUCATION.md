# Merge Many Streams

## Intuition

A k-way merge only ever needs the front element of each feed. That is k values in memory regardless of how much data flows through, which is the whole reason to merge instead of concatenating.

## Approach

1. Prime `heads` and `live` with one `Next()` per feed.
2. Each step, scan the live heads for the smallest, taking the earliest on a tie.
3. Append it and refill only that feed's slot.
4. Return when no feed is live.

## Solution

```go
func MergeAll(feeds ...Feed) []int {
	heads := make([]int, len(feeds))
	live := make([]bool, len(feeds))
	for i, f := range feeds {
		heads[i], live[i] = f.Next()
	}

	var out []int
	for {
		best := -1
		for i := range feeds {
			if !live[i] {
				continue
			}
			if best == -1 || heads[i] < heads[best] {
				best = i
			}
		}
		if best == -1 {
			return out
		}

		out = append(out, heads[best])
		heads[best], live[best] = feeds[best].Next()
	}
}
```

## Walkthrough

Strict `<` in the comparison keeps the first feed on ties, which is what makes the duplicate case deterministic. Only the consumed slot is refilled, so each element is pulled exactly once.

## Pitfalls

- Refilling every head each round, which drops values.
- Draining all feeds into one slice and sorting — O(total) memory, the problem being avoided.
- Treating a zero head as "drained": zero is a legal value, so track liveness separately.
