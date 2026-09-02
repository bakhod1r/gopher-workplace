# Compact By Equality

**Level:** middle  
**Topic:** 03-generics

## Context

A log collapses repeated lines, where "repeated" means the same message even if the timestamps differ.

## Task

Implement the stub(s) in [slicescompactfunc.go](slicescompactfunc.go):

1. Implement `Dedupe` using `slices.CompactFunc`.
2. Only adjacent equal elements collapse.
3. Leave the input untouched.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Dedupe(lines, sameMsg)
Output: runs collapsed
```

**Example 2:**

```
Input:  Dedupe([]int{1,2,1}, equal)
Output: []int{1,2,1}
```

**Example 3:**

```
Input:  Dedupe([]int{}, equal)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.CompactFunc`** | Like `Compact`, but equality is supplied rather than `==`. |
| 2 | **Custom equality** | Neither the element type nor its fields need to be comparable. |
| 3 | **Clone before mutating** | `slices.Clone` plus a nil guard is the standard pure-function wrapper. |

## Hint

Same shape as `Compact`, with the comparison passed in.

## Validate

```bash
make verify
```
