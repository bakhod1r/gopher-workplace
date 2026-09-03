# Shard Backup Upload

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

The nightly backup is split into shards that are uploaded to object storage in parallel. Operators need the full picture at the end of the run: which shards landed and which did not. Aborting the whole batch at the first rejected key would leave the backup half written with no record of what else would have failed.

## Task

Implement the exported function(s) in [shardbackup.go](shardbackup.go) so that:

1. Return a slice of errors the same length as `shards`, in input order.
2. Slot `i` is `nil` when `upload(shards[i])` succeeded, otherwise the error it returned.
3. Upload each shard in its own goroutine, joined with a `sync.WaitGroup`.
4. Attempt every shard: one failure must not cancel or skip the others.
5. A nil or empty shard list calls nothing and returns an empty slice.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  UploadShards([]Shard{{"a", 1}, {"b", 0}}, upload)
Output: [<nil> errEmptyShard]
```

**Example 2:**

```
Input:  UploadShards([]Shard{{"a", 0}, {"b", 1}, {"c", 0}}, upload)
Output: [errEmptyShard <nil> errEmptyShard]
```

**Example 3:**

```
Input:  UploadShards(nil, upload)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Per-index result slice** | Each goroutine owns exactly one slot, so no lock is needed and the race detector stays quiet. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before the launch, `defer wg.Done()` inside, `wg.Wait()` before reading the results. |
| 3 | **Loop-variable capture** | Pass `i` and the `Shard` value as goroutine parameters instead of closing over the loop variables. |
| 4 | **Partial failure** | Collecting every error beats returning the first one when the caller needs a full report. |

## Hint

Preallocate `make([]error, len(shards))`. Distinct indices of the same slice can be written from different goroutines without a mutex.

## Validate

```bash
make verify
```
