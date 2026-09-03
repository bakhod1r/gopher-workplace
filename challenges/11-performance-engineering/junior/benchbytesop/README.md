# The MB/s Column

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Call `b.SetBytes(n)` and the benchmark line grows an MB/s column. For anything that moves data — encoders, hashers, copies — throughput is the number that stays comparable when the input size changes.

## Task

Implement `ThroughputMBs` in [benchbytesop.go](benchbytesop.go):

1. Convert `totalBytes` moved in `elapsedNS` nanoseconds into megabytes per second.
2. Use 1 MB = 1,000,000 bytes and 1 s = 1,000,000,000 ns.
3. A non-positive `elapsedNS` returns `0`.

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
Input:  ThroughputMBs(1000000, 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Throughput vs latency** | MB/s stays flat across input sizes; ns/op does not. |
| 2 | **Decimal MB** | The benchmark tool uses 1e6, not 1<<20 — mixing them is a 5% error. |
| 3 | **Float division, integer inputs** | Convert before dividing or you truncate the answer to zero. |

## Hint

The two unit conversions cancel into a single ratio.

## Validate

```bash
make verify
```
