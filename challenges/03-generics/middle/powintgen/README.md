# Integer Power

**Level:** middle  
**Topic:** 03-generics

## Context

A hash mixer raises a prime to a configurable power. Multiplying in a loop is too slow for large exponents.

## Task

Implement the stub(s) in [powintgen.go](powintgen.go):

1. Implement `Pow` using repeated squaring, not a linear multiply loop.
2. Return `1` for `exp == 0` and the zero value for a negative `exp`.
3. The exponent is a plain `int`, not a `T`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Pow(2, 10)
Output: 1024
```

**Example 2:**

```
Input:  Pow(3, 0)
Output: 1
```

**Example 3:**

```
Input:  Pow(2, -1)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Mixed parameter types** | The base is generic; the exponent is always an `int`. |
| 2 | **Repeated squaring** | O(log exp) multiplications instead of O(exp). |
| 3 | **Identity element** | The accumulator starts at 1, as in any product. |

## Hint

Square the base each round and multiply into the result only on odd bits.

## Validate

```bash
make verify
```
