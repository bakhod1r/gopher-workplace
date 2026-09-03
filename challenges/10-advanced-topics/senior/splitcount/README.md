# Count Fields Without Splitting Anything

**Level:** senior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A CSV pre-pass reports each line's field count and payload size. It splits every line to count the pieces, then throws the pieces away — two allocations per line, on a file with a hundred million lines.

## Task

Fix the single planted bug in [splitcount.go](splitcount.go):

1. Return the number of sep-separated fields and the total byte size of those fields.
2. An empty input is 0, 0; a line with no separator is one field.
3. Fix the single bug so the function allocates nothing.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  CountFields([]byte("ab,c"), ',')
Output: 2, 3
```

_Explanation:_ Two fields, three payload bytes.

**Example 2:**

```
Input:  CountFields([]byte("a,,b"), ',')
Output: 3, 2
```

_Explanation:_ The empty middle field still counts.

**Example 3:**

```
Input:  CountFields(nil, ',')
Output: 0, 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Split allocates twice** | One copy for `string(line)` and one slice for the headers. |
| 2 | **Counting is not collecting** | Both answers are running totals; the pieces never need to exist. |
| 3 | **Separator arithmetic** | A line with n separators has n+1 fields. |

## Hint

You are asked for two numbers. How many of the pieces do you actually need to hold?

## Validate

```bash
make verify
```
