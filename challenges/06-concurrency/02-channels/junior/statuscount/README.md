# Count Status

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The access-log analyser scans a line stream for one HTTP status code and
counts the occurrences as they arrive, without ever holding the log in
memory.

## Task

Implement `CountStatus` in [statuscount.go](statuscount.go) so that:

1. It drains `lines` until the tailer closes the stream.
2. It counts lines equal to `want`.
3. An empty stream, or no matches, returns `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountStatus("200","500","200" | want "200")
Output: 2
```

**Example 2:**

```
Input:  CountStatus("200" | want "404")
Output: 0
```

**Example 3:**

```
Input:  CountStatus(empty | want "200")
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Streaming count** | Constant memory — the log is never buffered. |
| 2 | **String comparison** | `==` on strings compares contents, not pointers. |
| 3 | **`range` over a channel** | The loop ends when the tailer closes. |

## Hint

Count as lines arrive; there is no reason to collect the log into a slice
first.

## Validate

```bash
make verify
```
