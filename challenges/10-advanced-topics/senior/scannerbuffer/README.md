# The Scanner That Gave Up On A Long Line

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A log analyser works on every sample file and fails on the production log with "token too long". The offending line is a stack trace.

## Task

Fix the single planted bug in [scannerbuffer.go](scannerbuffer.go):

1. Return the length of the longest line in `r`.
2. Lines up to `maxLine` bytes must be accepted.
3. Fix the single bug; propagate any real read error.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  LongestLine(strings.NewReader("ab\ncdef"))
Output: 4, nil
```

**Example 2:**

```
Input:  a 200 KiB line
Output: its length, nil
```

_Explanation:_ Above the default limit.

**Example 3:**

```
Input:  a failing reader
Output: the error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Scanner.Buffer** | Sets the initial buffer and the maximum token size. |
| 2 | **The default 64 KiB cap** | Silent until a line exceeds it, then a hard error. |
| 3 | **Scanner.Err** | Where `bufio.ErrTooLong` surfaces — the loop just ends. |
| 4 | **Bounded by choice** | The limit protects against a hostile input; it should be explicit. |

## Hint

The loop ends early and `Err` explains why. Which knob was never turned?

## Validate

```bash
make verify
```
