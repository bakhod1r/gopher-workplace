# Greatest Common Divisor

**Level:** middle  
**Topic:** 03-generics

## Context

A layout engine reduces aspect ratios to their simplest form before storing them.

## Task

Implement the stub(s) in [gcdgen.go](gcdgen.go):

1. Implement `GCD` using Euclid's algorithm.
2. Negative inputs are handled by magnitude; the result is never negative.
3. `GCD(0, 0)` is `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  GCD(12, 18)
Output: 6
```

**Example 2:**

```
Input:  GCD(-12, 18)
Output: 6
```

**Example 3:**

```
Input:  GCD(5, 0)
Output: 5
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`%` needs an integer set** | Modulo is undefined for floats, so the constraint must exclude them. |
| 2 | **Euclid's loop** | `a, b = b, a%b` until `b` is zero — the classic two-line algorithm. |
| 3 | **Sign normalisation** | Taking magnitudes first keeps the result positive for any input signs. |

## Hint

Normalise the signs, then loop `a, b = b, a%b`.

## Validate

```bash
make verify
```
