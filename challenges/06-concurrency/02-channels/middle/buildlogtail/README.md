# Tail a Build Log

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

A build can emit hundreds of thousands of log lines, but the CI page only shows
the tail. The log tailer therefore has to consume the whole stream — so the
build's writer is never blocked — while holding at most `keep` lines in memory.
A buffered channel of capacity `keep` is the ring buffer.

## Task

Implement `TailBuildLog` in [buildlogtail.go](buildlogtail.go) so that:

1. It returns an empty, non-nil slice when `keep <= 0`, after draining the stream.
2. It creates a ring channel with capacity `keep`.
3. For each line it tries a non-blocking send into the ring; when the ring is full it receives the oldest line, discards it, and then sends.
4. When the stream closes it closes the ring, drains it into a slice, and returns the lines in stream order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TailBuildLog(log[a b c], keep=2)
Output: [b c]
```

**Example 2:**

```
Input:  TailBuildLog(log[a b], keep=5)
Output: [a b]
```

**Example 3:**

```
Input:  TailBuildLog(log[a b c d], keep=1)
Output: [d]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Capacity as a memory bound** | `make(chan string, keep)` is the whole retention policy. |
| 2 | **Drop-oldest** | Full ring: receive one, then send — the opposite of drop-newest. |
| 3 | **Closing a channel you own** | The ring is local, so closing it to `range` it is safe. |
| 4 | **FIFO ordering of a buffer** | A buffered channel hands values back in send order. |

## Hint

Even with `keep <= 0` you must still drain `lines`, otherwise the build's
writer stays blocked on a stream nobody is reading.

## Validate

```bash
make verify
```
