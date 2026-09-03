# Scaling That Divides First

**Level:** staff  
**Topic:** 03-generics

## Context

A billing job applies a percentage adjustment to line items in minor units. Small line items come out as zero and large ones are short by a few cents each, which across a month is a five-figure discrepancy.

## Task

Fix the single planted bug in [scaleroundbug.go](scaleroundbug.go):

1. Find and fix the single bug so the scaling multiplies before it divides, and rounds half away from zero.
2. Negative values must round away from zero too.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  ScaleAll([]int{7}, 300)
Output: []int{21}
```

**Example 2:**

```
Input:  ScaleAll([]int{250}, 15)
Output: []int{38}
```

**Example 3:**

```
Input:  ScaleAll([]int{-7}, 300)
Output: []int{-21}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Integer division truncates** | Dividing first throws away the remainder before the multiplication could have recovered it. |
| 2 | **Order of operations** | Doing the right steps in the wrong order is still a bug. |
| 3 | **Scale is a requirement** | A graded test asserts the result on millions of elements, so a defect that only shows past a threshold is caught. |

## Hint

Work out `7 * 300 / 100` both ways by hand.

## Validate

```bash
make verify
```
