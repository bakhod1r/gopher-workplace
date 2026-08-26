# Add Calculator

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A calculator object provides arithmetic operations as methods. This puzzle
highlights how a method differs from a standalone function: `calc.Add(2, 3)` vs
`Add(2, 3)`.

## Task

Implement `Add` on `Calculator` in [addcalc.go](addcalc.go):

1. Return `a + b`.
2. The struct is empty — it exists only to demonstrate method syntax.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Calculator{}.Add(2, 3)
Output: 5
```

**Example 2:**

```
Input:  Calculator{}.Add(-1, 1)
Output: 0
```

**Example 3:**

```
Input:  Calculator{}.Add(0, 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods vs functions** | A method is scoped to a type; a function is standalone. |
| 2 | **Empty struct** | `Calculator{}` costs 0 bytes — the struct is just a namespace. |
| 3 | **Value receiver** | Read-only; no state to mutate. |

## Hint

`return a + b` — trivial, but the point is the method syntax.

## Validate

```bash
make verify
```
