# Tail Build Log

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The CI build-log tailer replays a finished job's output newest-line-first
for the failure view. A producer goroutine walks the stored lines backwards
and the renderer collects whatever order it is given.

## Task

Implement `TailReverse` in [buildlog.go](buildlog.go) so that:

1. A goroutine sends the lines from last to first on an unbuffered channel.
2. It closes the channel when the stored log is exhausted.
3. The renderer collects them in arrival order; an empty log yields an empty, non-nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TailReverse(["a","b","c"])
Output: ["c" "b" "a"]
```

**Example 2:**

```
Input:  TailReverse(["x"])
Output: ["x"]
```

**Example 3:**

```
Input:  TailReverse(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Producer decides order** | The channel preserves send order; the producer chooses it. |
| 2 | **Backwards `for` loop** | `for i := len(lines)-1; i >= 0; i--`. |
| 3 | **Goroutine + unbuffered channel** | Sends interleave with the renderer's receives. |

## Hint

The channel does not reverse anything — the goroutine sends newest-first
and the channel keeps that order.

## Validate

```bash
make verify
```
