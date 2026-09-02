# Compact

**Level:** junior  
**Topic:** 03-generics

## Context

A sensor repeats its last reading while idle. The log keeps one entry per change.

## Task

Implement the stub(s) in [slicescompact.go](slicescompact.go):

1. Implement `Squash`, collapsing runs of equal neighbours into one element.
2. Leave the input untouched.
3. Note that non-adjacent duplicates are kept — `Compact` is not a deduplicator.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Squash([]int{1, 1, 2, 2, 1})
Output: []int{1, 2, 1}
```

**Example 2:**

```
Input:  Squash([]int{1, 2})
Output: []int{1, 2}
```

**Example 3:**

```
Input:  Squash([]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Compact`** | Removes only *consecutive* duplicates, and modifies the slice it is given. |
| 2 | **Cloning first** | Reused from earlier: clone before an in-place helper when the caller must keep its data. |
| 3 | **Compact is not Unique** | Deduplicating unsorted data still needs a set or a sort first. |

## Hint

`Compact` mutates its argument and returns the shortened slice — clone before calling.

## Validate

```bash
make verify
```
