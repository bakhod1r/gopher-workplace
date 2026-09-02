# Append If

**Level:** junior  
**Topic:** 03-generics

## Context

A query builder adds optional filter clauses. Each clause is appended only when its flag is set.

## Task

Implement the stub(s) in [appendifgen.go](appendifgen.go):

1. Implement `AppendIf`, appending `v` to `s` when `cond` is true.
2. Return `s` unchanged when `cond` is false.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AppendIf([]int{1}, 2, true)
Output: []int{1, 2}
```

**Example 2:**

```
Input:  AppendIf([]int{1}, 2, false)
Output: []int{1}
```

**Example 3:**

```
Input:  AppendIf([]string{}, "a", true)
Output: []string{"a"}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`append` returns a slice** | Reused from language basics: `append` may reallocate, so you must use its result. |
| 2 | **Type parameters** | `[T any]` declares a type parameter; the caller (or inference) picks `T`. |
| 3 | **Conditional builders** | Returning the slice makes calls chainable: `s = AppendIf(s, v, ok)`. |

## Hint

Two lines: a guard and an `append`.

## Validate

```bash
make verify
```
