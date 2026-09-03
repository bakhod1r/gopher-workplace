# Bounded Memory For An Unbounded Stream

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

"Keep the last thousand entries" is the standard answer to an unbounded stream: crash dumps, recent errors, a trace tail. A ring buffer does it with one array, one index and no allocation per entry — the arithmetic is a modulo, and the only subtlety is reading the entries back out in the right order.

## Task

Implement the five pieces in [ringbufwriter.go](ringbufwriter.go):

1. `New(n)` builds a ring of capacity `n`; a non-positive `n` holds nothing but still counts adds.
2. `Add` overwrites the oldest value once full and must not allocate.
3. `Len` is the current count, `Total` the lifetime count, and `Snapshot` returns the held values oldest-first in a fresh slice.

## Examples

**Example 1:**

```
Input:  New(3); Add(1..4); Snapshot()
Output: [2 3 4]
```

**Example 2:**

```
Input:  New(4); Add(1); Add(2); Snapshot()
Output: [1 2]
```

**Example 3:**

```
Input:  New(0); Add(1); Add(2); Total()
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The write index wraps** | `next = (next+1) % n` is the whole data structure. |
| 2 | **Oldest-first needs the offset** | Once wrapped, reading starts at `next`, not at index 0. |
| 3 | **Total outlives the contents** | Knowing how much you dropped is often as useful as what you kept. |

## Topics used again

Modular arithmetic, slices, methods on pointer receivers, defensive copies.

## Hint

While the ring is not yet full, `next` is also the length.

## Validate

```bash
make verify
```
