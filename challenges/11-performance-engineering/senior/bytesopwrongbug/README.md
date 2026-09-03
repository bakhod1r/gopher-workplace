# Megabytes That Are Not Megabytes

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

The throughput column disagrees with `go test -bench` by about five percent, consistently, on every machine. Nobody can find the missing work because there is none: two definitions of "megabyte" are in play, and only one of them is the tool's.

## Task

Fix the single planted bug in [bytesopwrongbug.go](bytesopwrongbug.go):

1. Find and fix the one bug so a million bytes in one second reports exactly 1 MB/s.
2. Halving the elapsed time must double the reported throughput.
3. The non-positive elapsed guard already works and must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  ThroughputMBs(1000000, 1000000000)
Output: 1
```

**Example 2:**

```
Input:  ThroughputMBs(1000000, 500000000)
Output: 2
```

**Example 3:**

```
Input:  ThroughputMBs(500, 1000)
Output: 500
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **MB versus MiB** | 1,000,000 and 1,048,576 differ by 4.86%, which is exactly the size of a plausible regression. |
| 2 | **Match the tool's units** | Your numbers are compared against `go test` output, so they must use its definition. |
| 3 | **Consistent small errors are unit errors** | A fixed percentage discrepancy is almost never real work. |

## Hint

One of the factors in the expression converts between two definitions of a megabyte.

## Validate

```bash
make verify
```
