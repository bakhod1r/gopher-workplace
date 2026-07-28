# Short-Circuit Nil Guard

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

`&&` evaluates left to right and short-circuits, so the nil check must come
FIRST: `p != nil && *p > 0`. With the operands swapped, `*p` runs on a nil
pointer before the guard, panicking.

## Task

Fix the guard order in [shortcircuit.go](shortcircuit.go).

Do **not** change the function signature or the tests.

## Examples

```go
ValueOr(nil, 5)  // => 5
ValueOr(&7, 5)   // => 7
ValueOr(&-1, 5)  // => 5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Short-circuit evaluation** | `&&` stops at the first false operand. |
| 2 | **Guard ordering** | Test `p != nil` before dereferencing. |
| 3 | **Nil safety** | `*p` on nil panics. |

## Hint

Put the nil check first: `if p != nil && *p > 0`.

## Validate

```bash
make verify
```
