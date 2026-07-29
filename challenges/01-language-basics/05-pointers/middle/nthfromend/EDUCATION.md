# The two-pointer gap technique

## Intuition

Keeping a fixed n-node gap between two pointers locates the n-th-from-end in one pass without counting length first.

## Approach

1. Advance a `lead` pointer `n` steps (nil if it overruns).
2. Move `lead` and a `trail` together until `lead` hits the end.
3. `trail` is then n-from-end.

## Solution

```go
type Node struct {
	Val  int
	Next *Node
}

func NthFromEnd(head *Node, n int) *Node {
	lead := head
	for i := 0; i < n; i++ {
		if lead == nil {
			return nil
		}
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

For `1..5, n=1`: lead advances 1 step to node 2, then lead and trail move together until lead is nil; trail lands on node 5.

## Pitfalls

- If advancing the lead n steps runs past the end, the list is too short.
- Both pointers move in lockstep after the gap is set.
