# Report The Best Run, Measure The Noise

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Benchmark noise is one-sided: nothing makes a machine faster than it really is, but a background process, a thermal throttle or a noisy neighbour all make it slower. That asymmetry is why the minimum of several runs is often the most honest single number, and why the max-over-min ratio is a good first check on whether the machine was quiet at all.

## Task

Implement both functions in [bestofn.go](bestofn.go):

1. `Best` returns the smallest sample and its index, earliest index on a tie.
2. `Best` on no samples returns `0, -1`.
3. `Spread` returns the largest sample divided by the smallest; no samples, or a smallest of zero, gives `0`.

## Examples

**Example 1:**

```
Input:  Best([5 2 9])
Output: 2, 1
```

**Example 2:**

```
Input:  Best([4 1 1 1])
Output: 1, 1
```

**Example 3:**

```
Input:  Spread([5 100 10])
Output: 20
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Noise only slows you down** | So the minimum estimates the true cost better than the mean. |
| 2 | **Spread as a quality gate** | A ratio far above 1 means the measurement environment, not the code, is what varied. |
| 3 | **The sentinel index** | Returning `-1` says "no samples" without a second return value. |

## Topics used again

Slices, `range` with an index, float comparison.

## Hint

One pass can track the minimum, its index, and the maximum together.

## Validate

```bash
make verify
```
