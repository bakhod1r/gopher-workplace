# Routing to Two Channels

## Intuition

Routing is a fan-out where the *value* picks the destination rather than
the scheduler. Because the split is data-dependent, you cannot size the
buffers tightly: worst case, every id goes one way. Both queues must also
be closed, since each drain loop ends only at its own close.

## Approach

1. Make `even` and `odd`, both with capacity `len(userIDs)`.
2. For each id, send to `even` if `id%2 == 0`, else to `odd`.
3. `close` both queues.
4. Drain each into a non-nil slice with `range`.
5. Return shard 0, then shard 1.

## Solution

```go
func SplitByShard(userIDs []int) ([]int, []int) {
	even := make(chan int, len(userIDs))
	odd := make(chan int, len(userIDs))
	for _, id := range userIDs {
		if id%2 == 0 {
			even <- id
		} else {
			odd <- id
		}
	}
	close(even)
	close(odd)

	shard0 := []int{}
	for id := range even {
		shard0 = append(shard0, id)
	}
	shard1 := []int{}
	for id := range odd {
		shard1 = append(shard1, id)
	}
	return shard0, shard1
}
```

## Walkthrough

For `[1, 2, 3, 4]`: `1` and `3` go to the odd queue, `2` and `4` to the
even queue. After closing, the two drains give `[2 4]` and `[1 3]`.

## Pitfalls

- Sizing the buffers at `len(userIDs)/2` blocks when the recipients are lopsided.
- Closing only one queue hangs the other drain loop.
- Draining the even queue before the odd one is closed still works here only because both closes happen first.
