# Hand Out Buffers Without A Lock On The Fast Path

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A pool behind a mutex is fine at four cores and is the top contended lock at sixty-four. The buffers are all the same size, and there are only ever a handful in flight.

## Task

Implement [ringpool.go](ringpool.go):

1. Return a buffer from the ring, emptied and ready to write.
2. Fall back to a fresh `p.size` buffer when every slot is empty.
3. A slot's buffer must be handed to exactly one caller — never two.
4. Correct under concurrent `Get` and `Put`.

Replace the stub body in [ringpool.go](ringpool.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NewBufPool(4, 64).Get()
Output: len 0, cap 64
```

**Example 2:**

```
Input:  Put then Get
Output: the same storage, emptied
```

**Example 3:**

```
Input:  4 slots filled, 4 Gets
Output: four distinct buffers
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Atomic claim by swap** | `Swap(nil)` both reads the slot and empties it, so only one caller wins. |
| 2 | **Read-modify-write must be one operation** | A load followed by a store lets two callers take one buffer. |
| 3 | **Bounded by construction** | A fixed ring cannot grow, whatever the load. |
| 4 | **Fallback over blocking** | An empty ring allocates rather than waiting. |

## Hint

Taking a slot has to be a single atomic operation. Which one both reads and clears?

## Validate

```bash
make verify
```
