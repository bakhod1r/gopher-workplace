# Average Reading

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The cold-chain monitor averages the temperature readings that arrived from
a shipping container during one window. A window with no readings has no
average — dividing by zero is not an answer.

## Task

Implement `AverageReading` in [tempavg.go](tempavg.go) so that:

1. It drains `readings`, tracking the running sum and the count.
2. It returns `sum / count` and `true` when at least one reading arrived.
3. It returns `0, false` for an empty window — no division by zero.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AverageReading(1, 2, 3)
Output: 2, true
```

**Example 2:**

```
Input:  AverageReading() // closed, empty
Output: 0, false
```

**Example 3:**

```
Input:  AverageReading(5)
Output: 5, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two accumulators** | A stream has no length, so count as you go. |
| 2 | **Guarding division** | Check `n == 0` before dividing. |
| 3 | **`float64(n)` conversion** | Go has no implicit int-to-float conversion. |

## Hint

Count while you sum — a channel has no `len` of its own once drained. Guard
the division when the count is zero.

## Validate

```bash
make verify
```
