# First Shard Error Cancels the Rest

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The analytics dashboard answers one question by scanning eight shards in parallel. If a single shard is unavailable the aggregate is worthless, so the other seven should stop scanning the moment the first failure is known. `errgroup.WithContext` is exactly this policy in one type: a `WaitGroup` that also owns a context it cancels on the first error.

## Task

Implement the exported function(s) in [shardquery.go](shardquery.go) so that:

1. It builds `group, groupCtx := errgroup.WithContext(ctx)`.
2. It starts one `group.Go` per shard, passing `groupCtx` — never the parent `ctx` — to the shard query.
3. It returns `group.Wait()`: nil if every shard succeeded, otherwise the first recorded error.
4. An error from one shard must cancel the others; errors that arrive afterwards are discarded.
5. An already-finished parent context reaches the shards as a cancelled `groupCtx`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  QueryAllShards(ctx, [])
Output: nil
```

**Example 2:**

```
Input:  QueryAllShards(ctx, [parked, broken, parked])
Output: errShardBroken  (both parked shards see context.Canceled)
```

**Example 3:**

```
Input:  QueryAllShards(cancelled ctx, [parked])
Output: context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`errgroup.WithContext`** | Returns a group plus a context it cancels as soon as any goroutine returns non-nil. |
| 2 | **First error wins** | `Wait` keeps only the first non-nil error; later ones are dropped. |
| 3 | **Passing the right context down** | Handing shards the parent `ctx` breaks cancellation — they must get `groupCtx`. |

## Hint

`errgroup.WithContext` hands back two things. The second one is the whole point — give it to the workers.

## Validate

```bash
make verify
```
