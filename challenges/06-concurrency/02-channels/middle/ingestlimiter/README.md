# Cap Indexing Concurrency

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

The crawler hands the indexer a batch of documents. Indexing them all at once
is the fastest way to knock the index host over, and doing them one at a time
wastes the batch window. The fix is a concurrency budget: at most `maxInFlight`
index calls in flight, enforced by a counting semaphore built from a buffered
channel.

## Task

Implement `IndexDocuments` in [ingestlimiter.go](ingestlimiter.go) so that:

1. It clamps `maxInFlight` to at least 1 and builds `slots := make(chan struct{}, maxInFlight)`.
2. It starts one goroutine per document; each acquires a slot (`slots <- struct{}{}`), calls `index`, and releases it (`<-slots`) with `defer`.
3. Results travel back on a channel; a closer goroutine waits for all workers and closes it.
4. The caller ranges the results channel into the returned map — non-nil even for an empty batch.

Every document is indexed exactly once.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IndexDocuments([a b c], maxInFlight=2, index)
Output: {a:10 b:20 c:30}, never more than 2 index calls at once
```

**Example 2:**

```
Input:  IndexDocuments([a], maxInFlight=8, index)
Output: {a:10}
```

**Example 3:**

```
Input:  IndexDocuments(nil, maxInFlight=4, index)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Counting semaphore** | Capacity is the budget; a full buffer blocks the next acquirer. |
| 2 | **`chan struct{}`** | A zero-width token — the value carries no data, only permission. |
| 3 | **Acquire/release symmetry** | `defer func(){ <-slots }()` releases even if `index` panics. |
| 4 | **Collecting results safely** | A results channel plus one closer avoids a mutex over the map. |

## Hint

Acquire the slot *inside* the goroutine, not before starting it — otherwise the
caller blocks and you have written a sequential loop with extra steps.

## Validate

```bash
make verify
```
