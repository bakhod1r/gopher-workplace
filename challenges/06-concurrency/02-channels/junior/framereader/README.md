# Read Frames

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The video transcoder's demuxer always asks for a fixed number of frame
sizes per group of pictures. When the upstream stream ends early the
remaining reads must not block — they come back as empty frames.

## Task

Implement `ReadFrames` in [framereader.go](framereader.go) so that:

1. It buffers `sizes`, closes the channel, then receives `len(sizes)+extra` times.
2. It returns everything received, in order.
3. Receives after the buffer is drained append `0`; a negative `extra` counts as `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ReadFrames([]int{1024, 512}, 1)
Output: [1024 512 0]
```

**Example 2:**

```
Input:  ReadFrames(nil, 2)
Output: [0 0]
```

**Example 3:**

```
Input:  ReadFrames([]int{5}, 0)
Output: [5]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closed-channel semantics** | Receives never block and yield the zero value forever. |
| 2 | **Buffered send before close** | Frames sent before `close` are still delivered. |
| 3 | **Zero value** | `0` for `int` — indistinguishable from a real `0` without comma-ok. |

## Hint

Closing does not discard buffered frames. Once they run out, `<-frames`
returns `0` immediately, as many times as the demuxer asks.

## Validate

```bash
make verify
```
