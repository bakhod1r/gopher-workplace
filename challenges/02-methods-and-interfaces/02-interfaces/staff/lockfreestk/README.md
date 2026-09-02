# Lock-Free Stack

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A hot free-list is contended by many goroutines. A mutex serialises them; a CAS loop lets them make progress independently.

## Task

Implement the stub(s) in [lockfreestk.go](lockfreestk.go):

1. Implement `Push` and `Pop` on `*Stack` with a compare-and-swap loop over `atomic.Pointer`.
2. Implement `Len`, which counts the nodes currently linked.
3. Constraint: `-race` clean under concurrent push and pop, no mutex, and no lost or duplicated elements.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Push(1), Push(2), Pop()
Output: 2 (LIFO)
```

**Example 2:**

```
Input:  Pop on an empty stack
Output: 0, false
```

**Example 3:**

```
Input:  1000 concurrent pushes and pops
Output: every value accounted for exactly once
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Compare-and-swap loops** | Read the head, prepare the new state, retry when the head moved. |
| 2 | **Lock-free progress** | A failed CAS means someone else made progress, so retrying is fair. |
| 3 | **Node allocation** | Each push allocates a node; the ABA problem is avoided because Go's GC keeps freed nodes alive. |

## Hint

`for { old := s.head.Load(); n.next = old; if s.head.CompareAndSwap(old, n) { return } }`.

## Validate

```bash
make verify
```
