# Divide And Remainder

**Level:** middle  
**Topic:** 03-generics

## Context

A pagination helper divides a total by a page size that comes from user input and may be zero.

## Task

Implement the stub(s) in [divmodgen.go](divmodgen.go):

1. Implement `DivMod`, returning the quotient, the remainder, and `true`.
2. Return zero values and `false` when `b` is zero — never panic.
3. Use Go's built-in semantics: the remainder takes the dividend's sign.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DivMod(7, 2)
Output: 3, 1, true
```

**Example 2:**

```
Input:  DivMod(-7, 2)
Output: -3, -1, true
```

**Example 3:**

```
Input:  DivMod(1, 0)
Output: 0, 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Integer division panics** | `a / 0` is a run-time panic for integers, unlike floats. |
| 2 | **Truncation towards zero** | Go truncates the quotient, so `-7 / 2` is `-3`, not `-4`. |
| 3 | **Remainder sign** | `a % b` takes the sign of `a` — different from a mathematical modulus. |

## Hint

The zero check is the whole puzzle; the arithmetic is built in.

## Validate

```bash
make verify
```
