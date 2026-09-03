# Store One Copy Of Each Repeated String

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A parser produces millions of strings, most of them repeats of a few hundred distinct values. Each repeat allocates its own copy, and the heap is full of identical short strings.

## Task

Implement [interncache.go](interncache.go):

1. Return a string with `b`'s contents, reusing the stored one when the bytes have been seen.
2. A repeat lookup must allocate nothing.
3. A stored string must own its bytes — the caller reuses its buffer.
4. The empty input returns `""` without storing anything.

Replace the stub body in [interncache.go](interncache.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  p.Intern([]byte("repeated")) twice
Output: the same string, one allocation
```

**Example 2:**

```
Input:  200 repeat lookups
Output: 0 allocations
```

**Example 3:**

```
Input:  26 batches through one reused buffer
Output: 26 distinct stored strings
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Borrow for the lookup, own for the store** | The read path may alias; the write path may not. |
| 2 | **Map keys must be immutable** | A stored key that changes strands its entry. |
| 3 | **Interning** | One canonical copy per distinct value, shared by every holder. |
| 4 | **Key and value are the same string** | Storing `owned` under itself keeps one allocation, not two. |

## Hint

Two paths: the hit borrows, the miss copies. Only one of them may allocate.

## Validate

```bash
make verify
```
