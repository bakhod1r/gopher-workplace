# Submission Ring

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

`io_uring` moves work between userspace and the kernel through a fixed-size
circular buffer: producers write at the tail, consumers read at the head, and
both indices wrap. No allocation, no growth — a full ring rejects the
submission.

## Task

Implement `Submit` and `Complete` on `*Ring` in [io_uring.go](io_uring.go):

1. `Submit(op)` returns `false` and changes nothing when `count == ringSize`.
   Otherwise it writes at `tail`, advances `tail` modulo `ringSize`, and
   increments `count`.
2. `Complete()` returns `(0, false)` when `count == 0`. Otherwise it reads at
   `head`, advances `head` modulo `ringSize`, decrements `count`, and returns
   the value with `true`.

**Constraint (staff):** after `New`, a `Submit`/`Complete` round trip must allocate nothing, and the cursors must stay correct across 100,000 wraps.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Submit(10), Submit(20), then two Complete()
Output: (10, true), (20, true)   — FIFO
```

**Example 2:**

```
Input:  Submit five times on a ring of size 4
Output: the fifth returns false; Len() stays 4
```

**Example 3:**

```
Input:  fill, drain 3, submit 3 more, drain all
Output: values come back in submission order across the wrap
```

_Explanation:_ the modulo is what makes slot reuse invisible to the caller.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Circular indices** | `(i + 1) % ringSize` is the whole of the wrap. |
| 2 | **Count disambiguates full from empty** | With only head and tail, `head == tail` means both — the explicit `count` avoids the classic ambiguity. |
| 3 | **Fixed-size array in a struct** | `[ringSize]int` is stored inline; the pointer receiver is what makes writes stick. |

## Hint

Check the guard **before** touching any index. A `Submit` that advances `tail`
and then discovers the ring is full has already corrupted the queue.

## Validate

```bash
make verify
```
