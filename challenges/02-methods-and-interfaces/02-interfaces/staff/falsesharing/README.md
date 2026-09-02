# False Sharing

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Per-worker counters sat in one array. Throughput collapsed on multi-core machines even though no two workers touched the same counter.

## Task

Implement the stub(s) in [falsesharing.go](falsesharing.go):

1. Implement `Inc` and `Total` on `*PaddedCounters`, where each counter occupies its own cache line.
2. Implement `Inc` and `Total` on `*PackedCounters`, the unpadded version, for comparison.
3. Constraint: both must be race-free and correct; the padded layout must place adjacent counters at least 64 bytes apart, which the test asserts.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  4 workers each Inc'ing their slot 1000 times
Output: Total 4000
```

**Example 2:**

```
Input:  the address distance between padded slots
Output: at least 64 bytes
```

**Example 3:**

```
Input:  Inc on an out-of-range slot
Output: ignored
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **False sharing** | Distinct variables on one cache line ping-pong that line between cores. |
| 2 | **Cache-line padding** | Pad each hot counter to its own line. |
| 3 | **unsafe.Sizeof** | Reused: verifying layout claims mechanically rather than by eye. |

## Hint

`struct { v atomic.Int64; _ [56]byte }` — 8 bytes of counter plus 56 bytes of padding fills a 64-byte line.

## Validate

```bash
make verify
```
