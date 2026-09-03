# The Buffered Writer Nobody Flushed

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

An export job writes a million rows and the file is short by a few kilobytes. The truncation is at the end, it varies run to run, and nothing reports an error.

## Task

Fix the single planted bug in [flushwriter.go](flushwriter.go):

1. Write each line followed by a newline through a buffered writer.
2. Fix the single bug so no output is lost.
3. Return any error, including one that only surfaces at the end.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  WriteAll(&buf, []string{"a","b"})
Output: "a\nb\n"
```

**Example 2:**

```
Input:  500 lines of 40 bytes
Output: 20500 bytes written
```

_Explanation:_ The final partial buffer must reach the writer.

**Example 3:**

```
Input:  a failing writer
Output: an error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Buffered writers hold output** | Bytes stay in the buffer until it fills or is flushed. |
| 2 | **Flush is where errors surface** | A write into the buffer succeeds even when the underlying writer will not. |
| 3 | **Silent truncation** | The missing bytes are always the last ones, which is why it looks random. |

## Hint

Every byte written so far is accounted for. What about the ones still in the buffer?

## Validate

```bash
make verify
```
