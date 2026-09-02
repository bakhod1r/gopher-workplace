# Split By Shard

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The chat fan-out router partitions recipient user ids across two delivery
shards by parity, pushing each id onto its shard's queue before the two
queues are handed to their senders.

## Task

Implement `SplitByShard` in [shardsplit.go](shardsplit.go) so that:

1. It sends each even user id to one queue and each odd id to the other.
2. It closes **both** queues, then drains each with `range`.
3. It returns shard 0 (even) first, shard 1 (odd) second, both in input order and both non-nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SplitByShard([]int{1, 2, 3, 4})
Output: [2 4], [1 3]
```

**Example 2:**

```
Input:  SplitByShard([]int{2})
Output: [2], []
```

**Example 3:**

```
Input:  SplitByShard(nil)
Output: [], []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Routing** | One producer, two destination queues chosen per id. |
| 2 | **Sizing both buffers** | Either queue might receive every id, so size both `len(userIDs)`. |
| 3 | **Closing every channel** | Each `range` needs its own queue closed. |

## Hint

Size both buffers at `len(userIDs)` — every recipient could land on one
shard — and close both before draining either.

## Validate

```bash
make verify
```
