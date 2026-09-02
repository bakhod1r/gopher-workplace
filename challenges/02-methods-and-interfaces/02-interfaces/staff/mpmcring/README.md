# MPMC Ring Buffer

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A hot queue between many producers and many consumers. A mutex serialises everything, so the slots carry their own sequence numbers instead.

## Task

Implement the stub(s) in [mpmcring.go](mpmcring.go):

1. Implement `Enqueue` and `Dequeue` on `*Ring` using per-slot sequence numbers and CAS on the head and tail.
2. `Enqueue` returns false when full; `Dequeue` returns false when empty — neither blocks.
3. Constraint: `-race` clean with many producers and consumers, no element lost or duplicated, and capacity must be a power of two so indexing is a mask.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  capacity 2; two Enqueues then a third
Output: true, true, false
```

**Example 2:**

```
Input:  Dequeue order
Output: FIFO
```

**Example 3:**

```
Input:  8 producers and 8 consumers
Output: every value delivered exactly once
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Vyukov MPMC queue** | Each slot's sequence number encodes whose turn it is. |
| 2 | **CAS on indices** | Producers and consumers advance independent counters. |
| 3 | **Power-of-two masking** | Reused: `i & (n-1)` replaces a modulo in the hot path. |

## Hint

Slot ready to write when `seq == pos`; ready to read when `seq == pos+1`. Publish by storing `pos+1` / `pos+n`.

## Validate

```bash
make verify
```
