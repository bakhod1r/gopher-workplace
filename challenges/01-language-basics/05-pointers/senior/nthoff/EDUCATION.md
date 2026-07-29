# Gap size in two-pointer windows

## Intuition

The lead pointer must be exactly n ahead; an inclusive bound makes it n+1 ahead and shifts the result.

## Approach

1. The lead pointer must advance exactly `n` steps, not `n+1`.
2. The bug's `i <= n` overshoots by one; use `i < n`.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func NthFromEnd(head *Node, n int) *Node {
	lead := head
	for i := 0; i < n; i++ {
		lead = lead.Next
	}
	trail := head
	for lead != nil {
		lead = lead.Next
		trail = trail.Next
	}
	return trail
}
```

## Walkthrough

For `n = 1`, `i <= n` moves lead 2 steps, landing the trailing pointer one node too early. `i < n` moves it exactly 1 step, so trail ends on the last node.

## Pitfalls

- `i <= n` advances n+1 times.
- The gap equals the number of lead-advance steps.
