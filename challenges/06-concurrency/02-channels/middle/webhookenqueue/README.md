# Shed Load on a Full Webhook Queue

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

Webhook deliveries are queued from inside the request that produced the event.
Blocking on a full queue would make a slow customer endpoint slow down our own
API, so the enqueue path is *non-blocking*: what fits is queued, what does not
fit is dropped to the dead-letter table and reported back.

## Task

Implement `EnqueueDeliveries` in [webhookenqueue.go](webhookenqueue.go) so that:

1. It offers each delivery to `queue` with a `select` that has a `default` arm, so it never blocks.
2. Ids that were sent go into `accepted`, in batch order.
3. Ids that hit a full queue go into `dropped`, in batch order — and it keeps offering the rest of the batch.
4. Both slices are non-nil, even when empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  EnqueueDeliveries(queue cap 2, [a b c])
Output: accepted [a b], dropped [c]
```

**Example 2:**

```
Input:  EnqueueDeliveries(queue cap 4, [a])
Output: accepted [a], dropped []
```

**Example 3:**

```
Input:  EnqueueDeliveries(unbuffered queue, [a])
Output: accepted [], dropped [a]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Non-blocking send** | `select { case ch <- v: default: }` sends only if it can proceed right now. |
| 2 | **Buffer capacity as a policy** | The queue's cap is how much lag you are willing to absorb. |
| 3 | **Send-only parameter** | `chan<- Delivery` proves this function never receives or closes. |
| 4 | **Load shedding** | Dropping with a record beats blocking the request path. |

## Hint

An unbuffered channel with nobody receiving is *always* full, so the `default`
arm takes every delivery — that is the third example, not a bug.

## Validate

```bash
make verify
```
