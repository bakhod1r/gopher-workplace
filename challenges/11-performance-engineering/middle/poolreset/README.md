# Pools That Do Not Grow Forever

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A pool with no size ceiling ratchets: one 100MB request puts a 100MB buffer back, and the pool holds it for the lifetime of the process while every subsequent 1KB request reuses that same monster. The fix is a cap on what `Put` accepts, and a `Get` that refuses to hand back something too small to be useful.

## Task

Implement all three methods in [poolreset.go](poolreset.go):

1. `Put` keeps buffers with capacity at most `MaxCap` and reports whether it kept the buffer; nil and oversized buffers are dropped.
2. `Get(n)` returns an empty buffer with capacity at least `n`, defaulting to 1024 for a non-positive `n`.
3. A pooled buffer that is too small for the request must not be returned; `Kept` reports how many buffers `Put` accepted.

## Examples

**Example 1:**

```
Input:  p.Put(make([]byte, 0, 1024))
Output: true
```

**Example 2:**

```
Input:  p.Put(make([]byte, 0, MaxCap+1))
Output: false
```

**Example 3:**

```
Input:  p.Put(cap 512); p.Get(8192)
Output: a buffer with cap at least 8192
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pools ratchet upward** | Without a ceiling the pool's memory is the largest request ever served. |
| 2 | **A pooled object may not fit** | `Get` has to check the capacity, not just accept whatever comes out. |
| 3 | **Reset before reuse** | Length zero on the way in and on the way out. |

## Topics used again

`sync.Pool`, mutexes for a counter, capacity checks, guards.

## Hint

`Kept` needs the mutex too — it reads state another goroutine may be writing.

## Validate

```bash
make verify
```
