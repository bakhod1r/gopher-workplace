# Compose Two Funcs

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Function composition builds a new function by chaining two — the returned
closure captures both `f` and `g`.

## Task

Implement `Compose` in [compose.go](compose.go) so the result computes `f(g(x))`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Compose(inc, double)(3)
Output: 7
```

**Example 2:**

```
Input:  applies g then f
Output: true
```

**Example 3:**

```
Input:  Compose(double, inc)(3)
Output: 8
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Capture two funcs** | Closure remembers `f` and `g`. |
| 2 | **Application order** | `f(g(x))`: g first, then f. |
| 3 | **Return a func** | Result type `func(int) int`. |

## Hint

`return func(x int) int { return f(g(x)) }`.

## Validate

```bash
make verify
```
