# Atomics And The Retry Loop

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A mutex around a single integer is a lot of machinery for one word. `atomic.Int64` does the same job with a single instruction — but only for operations the hardware supports directly. "Keep the largest value seen" is not one of them: load, compare and store is three steps, and the fix is a compare-and-swap loop that retries when someone else got there first.

## Task

Implement the four methods in [atomiccounter.go](atomiccounter.go):

1. `Add` increases the total by `delta` and returns the new value; `Total` reads it.
2. `Observe` keeps the largest value seen, using a compare-and-swap retry loop.
3. `Max` returns that maximum, `0` when nothing has been observed; the zero `Stats` must work and be race-free.

## Examples

**Example 1:**

```
Input:  Add(3), Add(-1)
Output: 3, then 2
```

**Example 2:**

```
Input:  Observe(5), Observe(2), Observe(4)
Output: Max is 5
```

**Example 3:**

```
Input:  100 goroutines observing values up to 10000
Output: Max is exactly 10000
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Atomic means one instruction** | Add and CAS are atomic; "load, decide, store" is not. |
| 2 | **The CAS retry loop** | Read, compute, swap-if-unchanged, and start over when the swap fails. |
| 3 | **Atomics are cheaper, not free** | The cache line still bounces between cores; they beat a mutex, they do not beat not sharing. |

## Topics used again

`sync/atomic`, compare-and-swap, loops, methods on pointer receivers.

## Hint

`for { cur := s.max.Load(); if v <= cur { return }; if s.max.CompareAndSwap(cur, v) { return } }`.

## Validate

```bash
make verify
```
