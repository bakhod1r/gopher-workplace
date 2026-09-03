# Averaging Without The Wild Ends

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Run a benchmark twenty times and one run will be garbage — a background build kicked in, a core got parked. The mean happily incorporates it. A trimmed mean throws away a fixed slice of each end first, which keeps the number stable without pretending outliers never happen.

## Task

Implement both functions in [latencytrim.go](latencytrim.go):

1. `TrimmedMean` sorts a copy, discards `floor(pct/100 * n)` samples from each end, and averages the rest.
2. `pct` is clamped into `[0,50)`, so at least one sample always survives; no samples gives `0`.
3. `Mean` is the plain average, and neither function may modify the input.

## Examples

**Example 1:**

```
Input:  TrimmedMean([1 2 3 4 100], 20)
Output: 3
```

**Example 2:**

```
Input:  Mean([1 2 3 4 100])
Output: 22
```

**Example 3:**

```
Input:  TrimmedMean([1 2 3], 99)
Output: the middle sample, not 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Trimming is symmetric** | Cutting only the top biases the estimate downwards. |
| 2 | **Clamp below 50%** | At 50% from each end there is nothing left to average. |
| 3 | **Trimming hides real tails** | Fine for benchmark noise, wrong for user-facing latency where the tail is the product. |

## Topics used again

`slices.Clone`, `slices.Sort`, `math.Floor`, slice bounds.

## Hint

After sorting, the surviving window is `s[k : len(s)-k]`.

## Validate

```bash
make verify
```
