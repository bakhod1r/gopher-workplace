# Cut The Tail Without Pinning It

## Intuition

Length is a view, not a fence. The array still holds the tail, and the collector follows the array, not your view of it. Clear the tail, then narrow the view.

## Approach

1. Clamp `n` low and high.
2. `clear(s[n:])` so the dropped pointers go to nil.
3. Return `s[:n]`.

## Solution

```go
// Node is one payload the slice points at.
type Node struct {
	ID int
}

// Truncate returns the first n elements of s, clearing the elements it
// drops so they no longer keep their payloads reachable.
//
// n is clamped into [0, len(s)]. The result reuses s's storage.
//
// Examples:
//
// 	Truncate([]*Node{{1}, {2}}, 1) => the first element only
func Truncate(s []*Node, n int) []*Node {
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	clear(s[n:])
	return s[:n]
}
```

## Walkthrough

With three nodes and n = 1: `clear(s[1:])` nils indices 1 and 2, then `s[:1]` is returned. Node 1 stays reachable, nodes 2 and 3 do not.

## Pitfalls

- Clearing after reslicing — `s[:n][n:]` is empty, so it clears nothing.
- Forgetting the clamp: `s[n:]` with n > len(s) panics.
