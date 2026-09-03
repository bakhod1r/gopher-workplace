# Longest Line, Bounded Working Set

**Level:** senior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A validator uses `bufio.Scanner` to find the longest record and dies with `token too long` on a 32 MiB line — then dies of memory when the buffer limit is raised.

## Task

Implement [maxline.go](maxline.go):

1. Return the byte length of the longest line, excluding the newline.
2. A trailing line without a newline still counts.
3. Hold only a fixed-size buffer: under 1 MiB total allocation for a 32 MiB line.

Replace the stub body in [maxline.go](maxline.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  MaxLine(strings.NewReader("ab\ncdef\ng"))
Output: 4, nil
```

_Explanation:_ "cdef" is the longest.

**Example 2:**

```
Input:  MaxLine(strings.NewReader("tail-without-newline"))
Output: 20, nil
```

_Explanation:_ The last line needs no terminator.

**Example 3:**

```
Input:  MaxLine(strings.NewReader("\n\n"))
Output: 0, nil
```

_Explanation:_ Two empty lines.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Counting instead of collecting** | The answer is a number, so the line never has to exist in memory. |
| 2 | **Reads do not align with lines** | The running count must survive across `Read` calls. |
| 3 | **EOF finalisation** | The last line is only complete once the stream ends. |

## Hint

You are asked for a length, not for the line.

## Validate

```bash
make verify
```
