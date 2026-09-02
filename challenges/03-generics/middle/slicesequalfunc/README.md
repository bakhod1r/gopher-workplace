# Equal With Custom Equality

**Level:** middle  
**Topic:** 03-generics

## Context

A test compares parsed rows against fixtures of a different shape, so `==` is unavailable.

## Task

Implement the stub(s) in [slicesequalfunc.go](slicesequalfunc.go):

1. Implement `SameRows` using `slices.EqualFunc`.
2. Different lengths are never equal; nil and empty are.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SameRows([]int{1}, []string{"1"}, matches)
Output: true
```

**Example 2:**

```
Input:  SameRows([]int{1}, []string{"2"}, matches)
Output: false
```

**Example 3:**

```
Input:  SameRows(nil, []string{}, matches)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.EqualFunc`** | Two element type parameters, bridged by the predicate. |
| 2 | **No `comparable` needed** | Equality comes from the function, so the elements can be anything. |
| 3 | **The `slices` package** | Generic slice helpers; the `Func` variants take a predicate or comparison. |

## Hint

One line — the stdlib already has the length check and the pairwise loop.

## Validate

```bash
make verify
```
