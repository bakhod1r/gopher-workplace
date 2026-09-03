# A Buffer Per Shard, Padded Apart

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

Every worker in a pool allocates a scratch buffer per call. Moving to one shared buffer created a race; moving to a `sync.Pool` helped, and the profile still shows cache-line ping-pong between the pool's per-shard state.

## Task

Implement [stripedbuffers.go](stripedbuffers.go):

1. Run `fn` on the shard's scratch buffer, reset to length 0 first.
2. Store the returned slice back and report its length.
3. Route `id` to a shard, including negative and out-of-range ids.
4. Hold only that shard's lock; safe for concurrent use.

Replace the stub body in [stripedbuffers.go](stripedbuffers.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s.With(0, appendOneByte)
Output: 1
```

**Example 2:**

```
Input:  ten calls on one shard
Output: 1 every time
```

_Explanation:_ The buffer is reset each call.

**Example 3:**

```
Input:  sizeof(stripe)
Output: 64
```

_Explanation:_ One shard per cache line.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Lock striping** | Independent shards mean independent locks. |
| 2 | **Padding computed from the fields** | `lineSize - Sizeof(mutex) - Sizeof(slice)` adapts if the struct changes. |
| 3 | **Reset before, store after** | `fn` may append past the capacity, so its result is the new buffer. |
| 4 | **Modulo with negative ids** | Go's `%` keeps the sign of the dividend. |

## Hint

Route, lock, reset, call, store, report.

## Validate

```bash
make verify
```
