# Greatest Common Divisor

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

The Euclidean algorithm repeatedly replaces the pair `(a, b)` with `(b, a%b)`
until `b` is 0. Signs must be normalized so the result is non-negative.

## Task

Implement `GCD(a, b)` (non-negative, `GCD(0,0)=0`).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  GCD(12, 8)
Output: 4
```

_Explanation:_ Euclid: 12%8=4, 8%4=0 -> 4

**Example 2:**

```
Input:  GCD(0, 9)
Output: 9
```

_Explanation:_ gcd with 0 is the other value

**Example 3:**

```
Input:  GCD(-12, 8)
Output: 4
```

_Explanation:_ absolute values used, result non-negative

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Modulo loop** | `a, b = b, a%b` until b==0. |
| 2 | **Sign handling** | Take absolute values first. |
| 3 | **Multiple assignment** | Update both in one statement. |

## Hint

Loop `for b != 0 { a, b = b, a%b }`; return `abs(a)`.

## Validate

```bash
make verify
```
