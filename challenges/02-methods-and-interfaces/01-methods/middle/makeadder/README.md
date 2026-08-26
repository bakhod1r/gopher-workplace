# Make Adder

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A method can act as a factory for closures. It captures the receiver's state
at the time the closure is created.

## Task

Implement `Adder` on `Number` in [makeadder.go](makeadder.go):

1. Return a function `func(x int) int`.
2. The function should return `n.Val + x`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  add := Number{5}.Adder(); add(3)
Output: 8
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closures** | Returning a function that accesses local variables (like `n`). |
| 2 | **Capture by value** | `n` is a value receiver, so its value is captured, not a pointer. |

## Hint

`return func(x int) int { return n.Val + x }`.

## Validate

```bash
make verify
```
