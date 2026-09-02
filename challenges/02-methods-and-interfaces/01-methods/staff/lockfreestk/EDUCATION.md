# Lock-Free Stack

## Intuition

A mutex makes other goroutines wait; a CAS loop makes them *retry*. The
difference matters under contention: with a lock, a descheduled holder stalls
everyone, while with CAS some goroutine always makes progress — that is what
"lock-free" formally means.

Push is the easy half of a Treiber stack, because a node is private until the
instant it becomes the head.

## Approach

1. Allocate the node once, outside the loop — it is not shared yet.
2. Each attempt: load the head, point the node at it, CAS.
3. Return on success; on failure the head moved, so start over with the new one.

## Solution

```go
func (s *Stack) Push(val int) {
	n := &node{val: val}
	for {
		old := s.head.Load()
		n.next = old
		if s.head.CompareAndSwap(old, n) {
			return
		}
	}
}
```

## Walkthrough

Two goroutines push at once. Both load the same `old`. One CAS wins and installs
its node; the other's CAS sees that the head is no longer `old` and fails. It
loops, loads the *new* head (the winner's node), repoints `n.next` at it and
tries again. Nothing is lost, which is why the test counts exactly 100 nodes.

Note that `n.next` is written before the CAS publishes `n` — the node is
unreachable until the swap succeeds, so that write needs no synchronization of
its own.

## Pitfalls

- **Loading `old` outside the loop.** After the first failure the CAS compares
  against a head that will never come back: an infinite spin.
- **Setting `n.next` outside the loop.** Same effect — on a retry the node still
  points at a stale head, and a successful CAS would then unlink every node
  pushed in between.
- **`if !CAS { return }`.** Silently drops pushes under contention; the count
  test is what catches it.
- **Assuming pop is symmetric.** It is not — popping introduces the ABA problem
  and the question of when a node can be freed. That is what hazard pointers
  (see the `hazardptr` puzzle) exist to answer.

## Why `-race` is not enough

The race detector will not flag a wrong-but-atomic algorithm; a dropped push is
a logic error, not a data race. The 100-goroutine count assertion is the test
that actually pins the behaviour down.
