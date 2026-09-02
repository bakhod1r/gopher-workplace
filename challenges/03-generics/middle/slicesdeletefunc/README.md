# Delete By Predicate

**Level:** middle  
**Topic:** 03-generics

## Context

A cleanup pass removes expired sessions from a slice the caller still needs in its original form.

## Task

Implement the stub(s) in [slicesdeletefunc.go](slicesdeletefunc.go):

1. Implement `Purge` using `slices.DeleteFunc`.
2. Leave the input untouched and return an empty (non-nil) slice for empty or nil input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Purge([]int{1,2,3}, isEven)
Output: []int{1,3}
```

**Example 2:**

```
Input:  Purge([]int{2}, isEven)
Output: []int{}
```

**Example 3:**

```
Input:  Purge(nil, isEven)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `slices` package** | Generic slice helpers; the `Func` variants take a predicate or comparison. |
| 2 | **In-place helpers** | Several `slices` functions rewrite their argument — clone first when the caller must keep it. |
| 3 | **Order is preserved** | `DeleteFunc` compacts survivors in place, keeping their relative order. |

## Hint

`slices.DeleteFunc` rewrites its argument — clone before handing it over.

## Validate

```bash
make verify
```
