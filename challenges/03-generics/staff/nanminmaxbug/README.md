# The Minimum That NaN Poisons

**Level:** staff  
**Topic:** 03-generics

## Context

A sensor pipeline reports the observed range of a metric. One dropped reading arrives as NaN and from then on both ends of the range render as `NaN` for the whole day.

## Task

Fix the single planted bug in [nanminmaxbug.go](nanminmaxbug.go):

1. Find and fix the single bug so NaN entries are skipped rather than adopted.
2. An all-NaN input must report `false`; integer instantiations must be unaffected.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  MinMaxSkipNaN([]float64{NaN, 3, 1, 2})
Output: 1, 3, true
```

**Example 2:**

```
Input:  MinMaxSkipNaN([]float64{NaN, NaN})
Output: 0, 0, false
```

**Example 3:**

```
Input:  MinMaxSkipNaN([]int{3, 1, 2})
Output: 1, 3, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **NaN is not ordered** | Every comparison involving NaN is false, so it neither wins nor loses a `<` test. |
| 2 | **Type sets** | A constraint names a *set* of types; the body must be correct for every member of it. |
| 3 | **Scale is a requirement** | A graded test asserts the result on millions of elements, so a defect that only shows past a threshold is caught. |

## Hint

What does `v < mn` do once `mn` is NaN?

## Validate

```bash
make verify
```
