# Metrics Fan In

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The metrics collector scrapes each node on its own goroutine, so samples
arrive on one channel per node. The aggregator needs a single stream to write
to storage, so the per-node channels are *fanned in*: one merged channel that
closes only after the last node has gone quiet.

## Task

Implement `MergeMetrics` in [metricsmerge.go](metricsmerge.go) so that:

1. It starts one forwarding goroutine per input stream, each pushing into a shared merged channel.
2. A separate goroutine waits for all forwarders and then closes the merged channel.
3. The caller drains the merged channel and returns the samples sorted ascending.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MergeMetrics(chan{1}, chan{2})
Output: []int{1, 2}
```

**Example 2:**

```
Input:  MergeMetrics(chan{3, 1}, chan{2})
Output: []int{1, 2, 3}
```

**Example 3:**

```
Input:  MergeMetrics()
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fan-in** | One forwarding goroutine per source; the merged channel has many writers. |
| 2 | **Closer goroutine** | `go func() { wg.Wait(); close(merged) }()` closes exactly once, after the last writer. |
| 3 | **Loop variable capture** | Pass the stream into the goroutine so each forwarder reads its own channel. |

## Hint

You cannot call `wg.Wait()` before draining — put the wait-and-close in its
own goroutine so the drain and the wait run at the same time.

## Validate

```bash
make verify
```
