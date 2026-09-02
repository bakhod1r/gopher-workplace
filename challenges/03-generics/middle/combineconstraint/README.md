# Combining Type Sets

**Level:** middle  
**Topic:** 03-generics

## Context

A table renderer accepts numeric and textual columns but must reject anything else, so the column type is checked at compile time.

## Task

Implement the stub(s) in [combineconstraint.go](combineconstraint.go):

1. Implement `Render`, formatting each element with `fmt.Sprint`.
2. Study `NumOrText`: it is the union of two constraint interfaces, not a new list of types.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Render([]int{1, 2})
Output: []string{"1", "2"}
```

**Example 2:**

```
Input:  Render([]string{"a"})
Output: []string{"a"}
```

**Example 3:**

```
Input:  Render([]float64{})
Output: []string{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Unions of constraints** | `interface{ Number | Text }` merges two type sets into one. |
| 2 | **Reusing constraint vocabulary** | Named constraints compose, so the sets stay defined in one place. |
| 3 | **Passing a type parameter to `any`** | `fmt.Sprint(v)` boxes the value — fine here, but it is a real conversion. |

## Hint

The constraint composes existing ones; the body is a plain formatting loop.

## Validate

```bash
make verify
```
