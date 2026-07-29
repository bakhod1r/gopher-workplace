# Euclidean GCD

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _control-flow_

## Context

The Euclidean algorithm repeatedly replaces `(a, b)` with `(b, a % b)` until `b` is zero; then `a` is the greatest common divisor.

## Task

Implement `GCD` in [gcdloop.go](gcdloop.go) using a loop.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  GCD(12, 8)
Output: 4
```

**Example 2:**

```
Input:  GCD(17, 5)
Output: 1
```

**Example 3:**

```
Input:  GCD(0, 9)
Output: 9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **for loop** | Iterate until `b == 0`. |
| 2 | **parallel assignment** | `a, b = b, a%b`. |
| 3 | **modulo** | `a % b` drives the reduction. |

## Hint

`for b != 0 { a, b = b, a%b }; return a`.

## Validate

```bash
make verify
```
