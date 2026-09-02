# Log Ingest Pipeline

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Wiring the ingest service end to end: a feeder goroutine streams raw lines off
the tailed file, a parse goroutine normalises each record, and the caller —
the last stage — keeps the error records for the alerting index. No shared
memory anywhere, only channels.

## Task

Implement `IngestPipeline` in [ingestpipeline.go](ingestpipeline.go) so that:

1. A feeder goroutine streams `lines` on a channel and closes it.
2. A second goroutine ranges over that channel, sends `parse(line)`, and closes its own output.
3. The calling goroutine ranges over the parsed channel and appends the records where `isError` holds, preserving input order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IngestPipeline([]string{"err disk", "info ok"}, upper, hasErrPrefix)
Output: []string{"ERR DISK"}
```

**Example 2:**

```
Input:  IngestPipeline([]string{"info ok"}, upper, hasErrPrefix)
Output: nil
```

**Example 3:**

```
Input:  IngestPipeline([]string{"err z", "info m", "err a"}, ...)
Output: []string{"ERR Z", "ERR A"}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stage composition** | Each stage's output channel is the next stage's input. |
| 2 | **Sequential ordering** | A single-channel chain preserves order — one record at a time. |
| 3 | **Draining the tail** | The caller is the last stage; it must drain so upstream goroutines can finish. |

## Hint

Build it bottom-up: feeder channel, then a parsed channel that ranges over
it, then range over the parsed channel in the caller.

## Validate

```bash
make verify
```
