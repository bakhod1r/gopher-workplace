# Log Parse Stage

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The log ingest service reads lines off a tailed file and pushes them through
parse → filter → index. Each step is a *stage*: it takes an input channel,
transforms every value, and hands the result to the next stage on a channel of
its own. The parsing work is injected so the stage stays testable.

## Task

Implement `ParseStage` in [logparse.go](logparse.go) so that:

1. It returns a new channel immediately.
2. A goroutine ranges over `lines` and sends `parse(line)` for each record.
3. When `lines` is closed and drained, the output channel is closed too.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  lines "warn disk", "info ok" with parse = strings.ToUpper
Output: "WARN DISK", "INFO OK" then closed
```

**Example 2:**

```
Input:  lines "boot" with parse = strings.ToUpper
Output: "BOOT" then closed
```

**Example 3:**

```
Input:  closed empty stream
Output: closed immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pipeline stage** | In-channel → transform → out-channel; stages compose by chaining. |
| 2 | **Close propagation** | Closing the input ends the `range`, which triggers closing the output. |
| 3 | **Injected work** | Passing `parse` in keeps the stage independent of the log format. |

## Hint

`for line := range lines` ends by itself when the input is closed — put
`defer close(out)` above it and shutdown chains down the pipeline for free.

## Validate

```bash
make verify
```
