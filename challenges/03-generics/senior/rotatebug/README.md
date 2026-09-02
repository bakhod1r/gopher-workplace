# Rotation That Panics On Negatives

**Level:** senior  
**Topic:** 03-generics

## Context

A carousel crashes as soon as the user drags it backwards past the first slide.

## Task

Fix the single planted bug in [rotatebug.go](rotatebug.go):

1. Find and fix the single bug so negative offsets rotate right.
2. Offsets larger than the length must still wrap.
3. Leave the input untouched.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Rotate([]int{1,2,3}, -1)
Output: []int{3,1,2}
```

**Example 2:**

```
Input:  Rotate([]int{1,2,3}, 4)
Output: []int{2,3,1}
```

**Example 3:**

```
Input:  Rotate([]int{}, 3)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sign of `%` in Go** | `-1 % 3` is `-1`, not `2` — the result takes the dividend's sign. |
| 2 | **Normalising into a range** | `((k % n) + n) % n` always lands in `[0, n)`. |
| 3 | **Panic versus wrong answer** | A negative slice index panics, which at least fails loudly. |

## Hint

What is `k` after the modulus when the caller passes `-1`?

## Validate

```bash
make verify
```
