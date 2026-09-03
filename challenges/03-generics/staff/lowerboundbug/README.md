# Binary Search That Overshoots The Duplicates

**Level:** staff  
**Topic:** 03-generics

## Context

A time-series API answers "give me everything from timestamp T onward" by binary-searching for the first matching sample. Whenever several samples share a timestamp, the response silently omits all of them.

## Task

Fix the single planted bug in [lowerboundbug.go](lowerboundbug.go):

1. Find and fix the single bug so the *first* element not less than `v` is found.
2. The search must stay logarithmic and keep returning `len(s)` for a value past the end.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  LowerBound([1,2,2,2,3], 2)
Output: 1
```

**Example 2:**

```
Input:  LowerBound([1,3], 9)
Output: 2
```

**Example 3:**

```
Input:  200000 queries over 4000000 elements
Output: well under the time budget
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Loop invariants** | The half-open invariant is "the answer lies in [lo, hi)"; the branch must preserve it. |
| 2 | **Lower versus upper bound** | One comparison operator is the whole difference between the two functions. |
| 3 | **Equal keys are the interesting case** | On a slice of distinct values both variants agree, so duplicates are the only witness. |

## Hint

Which side of the interval should an element *equal* to `v` end up on?

## Validate

```bash
make verify
```
