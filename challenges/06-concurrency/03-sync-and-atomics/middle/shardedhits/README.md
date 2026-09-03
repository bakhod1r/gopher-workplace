# Sharded Route Meter

**Level:** middle  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

The gateway records a hit for every request it proxies. At peak the process handles far more requests per second than a single mutex can hand out, and profiling shows workers queueing on the counter lock. The fix is sharding: split the map into independent shards so two requests on different routes rarely contend.

## Task

Implement the exported function(s) in [shardedhits.go](shardedhits.go) so that:

1. `Record` locks only the shard that owns the route and increments its counter.
2. `Count` returns a route's hits, or `0` when the route was never recorded.
3. `Total` sums every counter in every shard.
4. `Routes` returns every recorded route, sorted.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  m := NewMeter(4); m.Record("/orders"); m.Count("/orders")
Output: 1
```

**Example 2:**

```
Input:  NewMeter(4).Count("/unknown")
Output: 0
```

**Example 3:**

```
Input:  m.Record("/b"); m.Record("/a"); m.Routes()
Output: ["/a" "/b"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Lock sharding** | Independent locks over disjoint data cut contention without changing the result. |
| 2 | **Hash-to-shard** | The same route must always map to the same shard, or counts split in two. |
| 3 | **Aggregation across locks** | `Total` takes each shard lock in turn; it never holds two at once. |

## Hint

`shardFor` is already written. Every method that touches a shard's map must hold that shard's mutex — including the read-only ones.

## Validate

```bash
make verify
go test -race ./...
```
