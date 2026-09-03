# Deciding When A Number Is Bad News

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Benchmark numbers move a few percent between runs on the same code. A CI check that fails on any increase fails constantly and gets disabled within a week; one with a tolerance band fails only when something real happened. The band is the whole design.

## Task

Implement both functions in [regressionflag.go](regressionflag.go):

1. `Classify` returns `"regression"` above `+tolerance`, `"improvement"` below `-tolerance`, and `"noise"` inside the band — boundaries included.
2. A negative tolerance is treated as `0`.
3. `Failing` reports whether any change in the set is a regression.

## Examples

**Example 1:**

```
Input:  Classify(-20, 5)
Output: "improvement"
```

**Example 2:**

```
Input:  Classify(5, 5)
Output: "noise"
```

**Example 3:**

```
Input:  Failing([-10 2 30], 5)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Tolerance keeps the check alive** | A gate that cries wolf gets turned off, and then nothing is measured. |
| 2 | **Symmetric band, asymmetric consequences** | Improvements and regressions use the same threshold but not the same response. |
| 3 | **One bad benchmark fails the suite** | The check is an "any", not an average. |

## Topics used again

Float comparison, `switch`, loops over slices.

## Hint

`Failing` is one loop over `Classify`.

## Validate

```bash
make verify
```
