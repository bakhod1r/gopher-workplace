# The Popped Element That Never Left

## Intuition

Shortening a slice changes your view, not the memory. As long as the backing array is alive, every pointer still in it is alive too — so a stack that is reused forever retains everything it ever held.

## Approach

1. Handle the empty case.
2. Read the last element, write nil into its slot.
3. Return the element and `s[:i]`.

## Solution

```go
// Job is one queued unit of work.
type Job struct {
	ID  int
	Pad [1024]byte
}

// Pop removes and returns the last element of s.
//
// The shortened slice must not keep the popped job reachable: the element
// is still in the backing array until it is cleared.
//
// Examples:
//
// 	Pop([]*Job{a, b}) => b, a slice holding only a
func Pop(s []*Job) (*Job, []*Job) {
	if len(s) == 0 {
		return nil, s
	}
	i := len(s) - 1
	last := s[i]
	s[i] = nil
	return last, s[:i]
}
```

## Walkthrough

Popping eight jobs from an eight-slot stack leaves eight kilobytes reachable before the fix and none after it — the array is still there, but every slot is nil.

## Pitfalls

- Clearing after reslicing, which no longer has access to the slot.
- Assuming the collector can free part of an array; it frees allocations, not ranges.
