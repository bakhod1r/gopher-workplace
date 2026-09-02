# Index By Predicate

**Level:** junior  
**Topic:** 03-generics

## Context

A validation pass reports the position of the first bad reading so the user can jump to it.

## Task

Implement the stub(s) in [slicesindexfunc.go](slicesindexfunc.go):

1. Implement `FirstNegative` using `slices.IndexFunc`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstNegative([]int{1, -2, -3})
Output: 1
```

**Example 2:**

```
Input:  FirstNegative([]int{1, 2})
Output: -1
```

**Example 3:**

```
Input:  FirstNegative([]int{})
Output: -1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.IndexFunc`** | Takes a predicate instead of a value, for when equality is not the test. |
| 2 | **The `slices` package** | The stdlib ships generic slice helpers — prefer them over hand-rolled loops. |
| 3 | **Closures as arguments** | Reused from language basics: an inline `func(int) bool` is an ordinary value. |

## Hint

`slices.IndexFunc(nums, func(n int) bool { return n < 0 })`.

## Validate

```bash
make verify
```
