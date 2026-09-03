# Log Ingest Stages with Shutdown

**Level:** middle
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The log ingest service tails files, parses each line into a record, drops
everything that is not an error, and ships the survivors to the search index.
That is three stages joined by channels. The part that bites in production is
shutdown: when the service is asked to stop mid-batch, every stage has to
notice and return, or the process hangs on a send into a channel nobody reads.

## Task

Implement `IndexLogs` in [logindexstages.go](logindexstages.go) so that:

1. A reader goroutine streams `lines` onto a `raw` channel and closes it.
2. A parse goroutine ranges over `raw`, calls `parse`, skips records where the
   second return is false, forwards the rest on `parsed`, and closes `parsed`.
3. The caller is the last stage: it ranges over `parsed` and appends
   `index(rec)` in input order.
4. Every stage's send is a `select` against `done`, and the last stage stops
   appending once `done` is closed — so a shutdown never leaves a goroutine
   blocked.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IndexLogs(open done, ["err disk", "info ok", "err io"], parse, index)
Output: ["idx:disk", "idx:io"]
```

**Example 2:**

```
Input:  IndexLogs(open done, ["info ok", "warn slow"], parse, index)
Output: nil
```

**Example 3:**

```
Input:  IndexLogs(closed done, ["err disk", "err io"], parse, index)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stage composition** | Each stage owns its output channel and is the only one that closes it. |
| 2 | **Close propagation** | Closing `raw` ends the parse stage's `range`, which closes `parsed`, which ends the caller. |
| 3 | **or-done select** | `select { case ch <- v: case <-done: return }` is what makes a stage cancellable. |
| 4 | **Filtering stage** | Dropping a record is just `continue` — the stage sends fewer values than it receives. |

## Hint

Build it bottom-up and give each goroutine a `defer close(...)` for *its own*
output channel only. A stage never closes a channel it reads from.

## Validate

```bash
make verify
```
