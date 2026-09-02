# Is Sorted By

**Level:** middle  
**Topic:** 03-generics

## Context

A binary search assumes its input is sorted by name. A cheap precondition check catches misuse in tests before it becomes a mystery bug.

## Task

Implement the stub(s) in [slicesissortedfunc.go](slicesissortedfunc.go):

1. Implement `ByName` using `slices.IsSortedFunc`.
2. Equal neighbours are allowed; empty and single-element slices are sorted.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ByName(sorted)
Output: true
```

**Example 2:**

```
Input:  ByName(unsorted)
Output: false
```

**Example 3:**

```
Input:  ByName(nil)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.IsSortedFunc`** | Checks adjacent pairs with your comparison. |
| 2 | **Precondition checks** | Cheap in tests, and they turn silent wrongness into a clear failure. |
| 3 | **Non-decreasing** | Equal neighbours must not fail the check. |

## Hint

One call — the comparison is the same one you would pass to the sort.

## Validate

```bash
make verify
```
