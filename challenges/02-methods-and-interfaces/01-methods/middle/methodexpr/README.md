# Method Expression

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A generic dispatcher calls methods by name. Method expressions let you treat
`Type.Method` as a regular function where the receiver becomes the first
argument.

## Task

Implement `CallExpr` in [methodexpr.go](methodexpr.go):

1. Call `fn` with `a` and `n`.
2. `fn` is a method expression: `func(Adder, int) int`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CallExpr(Adder.Add, Adder{10}, 5)
Output: 15
```

**Example 2:**

```
Input:  CallExpr(Adder.Add, Adder{0}, 0)
Output: 0
```

**Example 3:**

```
Input:  CallExpr(Adder.Add, Adder{-3}, 7)
Output: 4
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method expressions** | `Adder.Add` has type `func(Adder, int) int` — receiver is first arg. |
| 2 | **Method values vs expressions** | Value: `a.Add` binds receiver. Expression: `Adder.Add` does not. |

## Hint

`return fn(a, n)` — the method expression takes the receiver as its first
parameter.

## Validate

```bash
make verify
```
