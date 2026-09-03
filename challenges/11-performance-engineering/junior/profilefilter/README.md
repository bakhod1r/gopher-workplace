# Hiding The Noise

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A real profile has thousands of nodes and the long tail is all noise: functions with one or two samples that will never be worth optimising. `pprof -nodefraction=0.05` throws them away so the graph stays readable. That flag is a threshold filter.

## Task

Implement `Filter` in [profilefilter.go](profilefilter.go):

1. Keep rows whose value is at least `minPct` percent of `total`.
2. Preserve the input order and leave the input slice untouched.
3. A non-positive `total` keeps nothing and returns an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  Filter([{a 50} {b 1} {c 20}], 100, 5)
Output: [{a 50} {c 20}]
```

**Example 2:**

```
Input:  Filter([{a 5}], 100, 5)
Output: [{a 5}]
```

**Example 3:**

```
Input:  Filter([{a 50}], 0, 5)
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Thresholding is reading strategy** | The tail of a profile is sampling noise, not work you can remove. |
| 2 | **Inclusive boundary** | "At least 5%" must keep exactly 5%, or the flag silently drops a real node. |
| 3 | **Filtering into a new slice** | Appending to a fresh slice keeps the caller's data and the order intact. |

## Topics used again

Slices, float comparison, `append` with a capacity hint.

## Hint

Compare `float64(e.Value) >= float64(total)*minPct/100` and append the survivors.

## Validate

```bash
make verify
```
