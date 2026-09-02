# Lock-Free Stack

## Intuition

A CAS loop replaces mutual exclusion with optimistic retry: read the head, build the new state against it, and swap only if nothing changed. A failed swap means another goroutine succeeded — the system as a whole always progresses.

## Approach

1. `Push` allocates a node, then loops: load the head, point `n.next` at it, CAS the head to `n`.
2. `Pop` loops: load the head, return `false` when nil, CAS the head to `old.next`.
3. Only on a successful CAS is the value returned.
4. `Len` walks the chain — a best-effort count, valid only when quiescent.

## Solution

```go
func (s *Stack) Push(v int) {
	n := &node{value: v}
	for {
		old := s.head.Load()
		n.next = old
		if s.head.CompareAndSwap(old, n) {
			return
		}
	}
}

func (s *Stack) Pop() (int, bool) {
	for {
		old := s.head.Load()
		if old == nil {
			return 0, false
		}
		if s.head.CompareAndSwap(old, old.next) {
			return old.value, true
		}
	}
}
```

## Walkthrough

Setting `n.next` *inside* the loop is essential: after a failed CAS the head has moved, so the previously recorded `next` is stale and re-swapping would drop every node pushed in between.

## Pitfalls

- Setting `n.next` once before the loop, which silently loses concurrently pushed nodes.
- Reading `old.value` before the CAS succeeds and returning it anyway — another popper may have taken that node.
- Assuming this is ABA-safe in a language without GC; here the collector keeps a popped node alive while anyone still references it.
