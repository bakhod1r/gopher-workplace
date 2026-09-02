# Contains By Predicate

**Level:** junior  
**Topic:** 03-generics

## Context

A sweeper only runs when at least one cache entry has expired, so the check must be cheap and short-circuiting.

## Task

Implement the stub(s) in [slicescontainsfunc.go](slicescontainsfunc.go):

1. Implement `AnyExpired` using `slices.ContainsFunc`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AnyExpired([{a 5} {b 0}])
Output: true
```

**Example 2:**

```
Input:  AnyExpired([{a 5}])
Output: false
```

**Example 3:**

```
Input:  AnyExpired([])
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.ContainsFunc`** | The predicate form of `Contains`, for tests other than equality. |
| 2 | **Short-circuiting** | It returns at the first match, like the `Any` you wrote by hand. |
| 3 | **The `slices` package** | The stdlib ships generic slice helpers — prefer them over hand-rolled loops. |

## Hint

Same shape as `IndexFunc`, but you only need the yes/no answer.

## Validate

```bash
make verify
```
