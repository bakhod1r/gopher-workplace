# Combined Reading

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

A building-automation dashboard shows a combined figure from a paired
temperature and humidity sensor. Both readings arrive on their own channels
and the order they land in is not fixed.

## Task

Implement `CombinedReading` in [sensorpair.go](sensorpair.go) so that:

1. It uses a `select` with one case per sensor channel.
2. It performs exactly two receives in total, one from each channel.
3. It returns the sum of the two readings.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CombinedReading(21, 40)
Output: 61
```

**Example 2:**

```
Input:  CombinedReading(0, 55)
Output: 55
```

**Example 3:**

```
Input:  CombinedReading(-5, 5)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`select`** | Waits on several channel operations; runs the first that is ready. |
| 2 | **Random choice** | If several cases are ready, `select` picks one at random. |
| 3 | **Order independence** | Summing makes the result the same whichever case fires first. |

## Hint

Loop twice around a two-case `select`. Each sensor sends exactly one
reading, so both get consumed.

## Validate

```bash
make verify
```
