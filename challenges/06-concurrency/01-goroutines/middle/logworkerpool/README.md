# Log Shipper Worker Pool

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

The log shipper drains a buffer of lines into the ingest endpoint. One goroutine per line would open one connection per line during a burst; instead a fixed pool of workers pulls line indices off a channel. The pool size is the connection budget, and it does not move when the burst does.

## Task

Implement the exported function(s) in [logworkerpool.go](logworkerpool.go) so that:

1. Return a slice of errors the same length as `lines`, in input order.
2. Start exactly `workers` goroutines, never one per line; treat `workers <= 0` as `1`.
3. Feed line indices to the workers over a channel and close it when the last index is sent.
4. Each worker ranges over the channel and writes its result into that line's slot.
5. `wg.Wait()` before returning, so no worker is still shipping when the caller reads the results.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ShipLines([]string{"ok", "bad-1", "ok"}, 2, ship)
Output: [<nil> errRejected <nil>]
```

**Example 2:**

```
Input:  ShipLines([]string{"ok", "bad-1"}, 0, ship)
Output: [<nil> errRejected]
```

**Example 3:**

```
Input:  ShipLines(nil, 4, ship)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Worker pool** | N long-lived goroutines consuming a job channel, instead of one short-lived goroutine per job. |
| 2 | **Channel ownership** | The sender owns the channel and is the only one allowed to `close` it. |
| 3 | **`for i := range jobs`** | The range loop ends exactly when the channel is closed and drained — that is the shutdown signal. |
| 4 | **Index as the job** | Sending the index, not the line, lets the worker write straight into the right result slot. |

## Hint

Send indices from the parent, `close(jobs)` after the send loop, then `wg.Wait()`. Closing before the workers finish is fine — closing is not the same as stopping.

## Validate

```bash
make verify
```
