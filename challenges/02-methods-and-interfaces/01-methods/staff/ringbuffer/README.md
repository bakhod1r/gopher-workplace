# Ring Buffer

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A fixed-capacity FIFO with no allocation after construction: writers advance a
tail, readers advance a head, and both wrap. Because `head == tail` describes
both a full and an empty buffer, an explicit `size` field settles the ambiguity.

## Task

Implement `Push` and `Pop` on `*RingBuffer` in [ringbuffer.go](ringbuffer.go):

1. `Push` returns `ErrFull` and changes nothing when `size == len(r.data)`.
   Otherwise write at `tail`, advance `tail` modulo capacity, increment `size`.
2. `Pop` returns `ErrEmpty` when `size == 0`. Otherwise read at `head`, advance
   `head` modulo capacity, decrement `size`, and return the value.

**Constraint (staff):** steady-state `Push`/`Pop` must allocate nothing, and a rejected `Push` must leave the buffer byte-for-byte unchanged.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  New(2); Push(1), Push(2), Push(3)
Output: nil, nil, ErrFull — and Len() stays 2
```

**Example 2:**

```
Input:  Push(1), Push(2), Pop(), Pop(), Pop()
Output: 1, 2, then ErrEmpty
```

**Example 3:**

```
Input:  New(3): fill, pop twice, push twice, drain
Output: 3, 4, 5 — order survives the wrap
```

_Explanation:_ freed slots are reused from the front of the array.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sentinel errors** | `errors.Is(err, ErrFull)` compares identity, so the errors are package-level vars. |
| 2 | **Modular cursors** | `(i + 1) % len(r.data)` is the wrap. |
| 3 | **Guard before mutate** | A rejected operation must leave every field untouched. |

## Hint

Check the guard first, always. A `Push` that advances `tail` and *then* notices
the buffer was full has permanently desynchronized the cursors.

## Validate

```bash
make verify
```
