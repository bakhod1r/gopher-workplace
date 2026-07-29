# Apply N Times

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Capturing both a function and a count produces a closure that self-iterates —
function exponentiation.

## Task

Implement `Repeat` in [repeatfn.go](repeatfn.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Repeat(inc, 3)(0)
Output: 3
```

**Example 2:**

```
Input:  Repeat(double, 2)(1)
Output: 4
```

**Example 3:**

```
Input:  Repeat(inc, 0)(5)
Output: 5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Capture f and n** | Both are remembered by the closure. |
| 2 | **Loop application** | Apply f n times inside. |
| 3 | **Identity at 0** | No applications returns the input. |

## Hint

`return func(x int) int { for i := 0; i < n; i++ { x = f(x) }; return x }`.

## Validate

```bash
make verify
```
