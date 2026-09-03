# Log Encoder Buffer Pool

**Level:** junior
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A log shipper renders millions of lines a minute. Allocating a fresh scratch buffer per line burns CPU in the garbage collector, so the encoder keeps a `sync.Pool` of byte slices that many goroutines share.

## Task

Implement the stubbed functions in [bufferpool.go](bufferpool.go) so that:

1. `NewEncoder` builds a pool whose `New` returns an empty `[]byte` with spare capacity.
2. `Encode` joins the fields with `|`, using a buffer from the pool.
3. The buffer is reset before use and returned to the pool afterwards.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewEncoder().Encode([]string{"warn", "disk full"})
Output: "warn|disk full"
```

**Example 2:**

```
Input:  NewEncoder().Encode([]string{"solo"})
Output: "solo"
```

**Example 3:**

```
Input:  NewEncoder().Encode(nil)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `sync.Pool` | A free list of reusable objects. `Get` returns one (or calls `New`), `Put` hands it back. |
| 2 | Reset before use | A pooled object carries the previous user's data. `buf[:0]` keeps the capacity, drops the contents. |
| 3 | Pool is not a cache | The GC may drop pooled objects at any time — `New` must always be able to build a fresh one. |

## Hint

`e.pool.Get().([]byte)` gives you a slice with unknown length. Slice it to `[:0]` before appending, and `Put` it back once you have built the string.

## Validate

```bash
make verify
```
